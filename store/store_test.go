package store_test

import (
	"sync"
	"testing"

	"github.com/no-src/nssession/store"
	"github.com/no-src/nssession/store/boltdb"
	"github.com/no-src/nssession/store/buntdb"
	"github.com/no-src/nssession/store/etcd"
	"github.com/no-src/nssession/store/fastcache"
	"github.com/no-src/nssession/store/memcached"
	"github.com/no-src/nssession/store/memory"
	"github.com/no-src/nssession/store/redis"
	"github.com/no-src/nssession/store/redis_cluster"
)

func TestStore(t *testing.T) {
	testCases := []struct {
		driver store.DriverName
		conn   string
	}{
		{memory.Driver, "memory:"},
		{buntdb.Driver, "buntdb://:memory:"},
		{redis.Driver, "redis://127.0.0.1:6379"},
		{redis_cluster.Driver, "redis-cluster://127.0.0.1:7001?addr=127.0.0.1:7002&addr=127.0.0.1:7003"},
		{etcd.Driver, "etcd://127.0.0.1:2379?dial_timeout=5s"},
		{boltdb.Driver, "boltdb://boltdb.db"},
		{memcached.Driver, "memcached://127.0.0.1:11211"},
		{fastcache.Driver, "fastcache://?max_bytes=50mib"},
	}

	for _, tc := range testCases {
		t.Run(string(tc.driver), func(t *testing.T) {
			s := store.NewStore(tc.driver)
			conn := tc.conn
			wg := sync.WaitGroup{}
			for i := 0; i < 5; i++ {
				wg.Add(1)
				go func() {
					if _, err := s.NewCache(conn); err != nil {
						t.Errorf("get cache component error by concurrent, err=%v", err)
					}
					wg.Done()
				}()
			}
			if _, err := s.NewCache(conn); err != nil {
				t.Errorf("get cache component error, err=%v", err)
			}
			wg.Wait()
		})
	}
}

func TestStore_ReturnError(t *testing.T) {
	testCases := []struct {
		driver store.DriverName
		conn   string
	}{
		{buntdb.Driver, "memory:"},
		{memory.Driver, ""},
		{etcd.Driver, "etcd://127.0.0.1:2379?dial_timeout=5z"},
	}

	for _, tc := range testCases {
		t.Run(string(tc.driver), func(t *testing.T) {
			s := store.NewStore(tc.driver)
			_, err := s.NewCache(tc.conn)
			if err == nil {
				t.Errorf("expect to get an error, but get nil")
			}
		})
	}
}
