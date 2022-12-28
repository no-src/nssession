package nssession

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/no-src/nssession/store"
	"github.com/no-src/nssession/store/buntdb"
	"github.com/no-src/nssession/store/etcd"
	"github.com/no-src/nssession/store/memory"
	"github.com/no-src/nssession/store/redis"
)

var (
	memoryConfig = &Config{
		Connection:    "memory:",
		Expiration:    time.Hour,
		CookieName:    DefaultCookieName,
		SessionPrefix: DefaultSessionPrefix,
		Store:         store.NewStore(memory.Driver),
	}

	buntDBConfig = &Config{
		Connection:    "buntdb://:memory:",
		Expiration:    time.Hour,
		CookieName:    DefaultCookieName,
		SessionPrefix: DefaultSessionPrefix,
		Store:         store.NewStore(buntdb.Driver),
	}

	redisConfig = &Config{
		Connection:    "redis://127.0.0.1:6379",
		Expiration:    time.Hour,
		CookieName:    DefaultCookieName,
		SessionPrefix: DefaultSessionPrefix,
		Store:         store.NewStore(redis.Driver),
	}

	etcdConfig = &Config{
		Connection:    "etcd://127.0.0.1:2379?dial_timeout=5s",
		Expiration:    time.Hour,
		CookieName:    DefaultCookieName,
		SessionPrefix: DefaultSessionPrefix,
		Store:         store.NewStore(etcd.Driver),
	}
)

func TestNSSession_Memory(t *testing.T) {
	testNSSession(t, memoryConfig)
}

func TestNSSession_BuntDB(t *testing.T) {
	testNSSession(t, buntDBConfig)
}

func TestNSSession_Redis(t *testing.T) {
	testNSSession(t, redisConfig)
}

func TestNSSession_Etcd(t *testing.T) {
	testNSSession(t, etcdConfig)
}

func testNSSession(t *testing.T, c *Config) {
	InitDefaultConfig(c)
	session, err := Default(createTestRequestAndWriter())
	if err != nil {
		t.Errorf("get session component error, err=%v", err)
		return
	}
	k := "hello"
	v := "world"
	var actual string

	err = session.Set(k, v)
	if err != nil {
		t.Errorf("set session data error, err=%v", err)
		return
	}

	err = session.Get(k, &actual)
	if err != nil {
		t.Errorf("get session data error, err=%v", err)
		return
	}
	if v != actual {
		t.Errorf("expect to get value %s, but get %v", v, actual)
		return
	}

	err = session.Remove(k)
	if err != nil {
		t.Errorf("remove session data error, err=%v", err)
		return
	}

	err = session.Get(k, &actual)
	if !errors.Is(err, ErrNil) {
		t.Errorf("expect to get error %v, but get %v", ErrNil, err)
		return
	}

	err = session.Set(k, v)
	if err != nil {
		t.Errorf("set session data error, err=%v", err)
		return
	}

	err = session.Get(k, &actual)
	if err != nil {
		t.Errorf("get session data error, err=%v", err)
		return
	}
	if v != actual {
		t.Errorf("expect to get value %s, but get %v", v, actual)
		return
	}

	err = session.Clear()
	if err != nil {
		t.Errorf("clear session data error, err=%v", err)
		return
	}

	err = session.Get(k, &actual)
	if !errors.Is(err, ErrNil) {
		t.Errorf("expect to get error %v, but get %v", ErrNil, err)
		return
	}

	err = session.Remove(k)
	if err != nil {
		t.Errorf("remove session data error, err=%v", err)
		return
	}
}

func TestNew_WithNilConfig(t *testing.T) {
	req, writer := createTestRequestAndWriter()
	_, err := New(nil, req, writer)
	if !errors.Is(err, errNilConfig) {
		t.Errorf("expect to get error %v, but get %v", errNilConfig, err)
	}
}

func TestNew_WithUnsupportedStoreDriver(t *testing.T) {
	c := &Config{
		Connection:    "memory:",
		Expiration:    time.Hour,
		CookieName:    DefaultCookieName,
		SessionPrefix: DefaultSessionPrefix,
		Store:         store.NewStore(redis.Driver),
	}
	req, writer := createTestRequestAndWriter()
	_, err := New(c, req, writer)
	if err == nil {
		t.Errorf("expect to get an error, but get nil")
	}
}

func TestNSSession_WithExistSessionID(t *testing.T) {
	InitDefaultConfig(memoryConfig)

	sessionID := "abcdefg"
	req, writer := createTestRequestAndWriter()
	req.AddCookie(&http.Cookie{
		Name:  memoryConfig.CookieName,
		Value: sessionID,
	})
	session, err := Default(req, writer)
	if err != nil {
		t.Errorf("get session component error, err=%v", err)
		return
	}
	if session.ID() != sessionID {
		t.Errorf("expect to get session id %s, but get %s", sessionID, session.ID())
	}
}

func createTestRequestAndWriter() (req *http.Request, writer http.ResponseWriter) {
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	writer = httptest.NewRecorder()
	return req, writer
}
