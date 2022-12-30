package nssession

import (
	"errors"
	"net/http"
	"time"

	"github.com/no-src/nssession/store"
)

var (
	defaultConfig *Config

	errNilConfig = errors.New("config is nil")
	errNilStore  = errors.New("store is nil")

	// DefaultCookieName the default cookie name for session
	DefaultCookieName = "ns-session-id"
	// DefaultSessionPrefix the default session prefix for session store
	DefaultSessionPrefix = "session"
)

// Config the configuration of session
type Config struct {
	// Connection the connection string of the session
	Connection string
	// Expiration the expiration time of the session
	Expiration time.Duration
	// SessionPrefix the session prefix for session store
	SessionPrefix string
	// Store the session store component
	Store store.Store
	// Cookie the settings of the cookie
	Cookie Cookie
}

// Cookie the settings of the cookie
type Cookie struct {
	// Name the cookie name for session
	Name string
	// Path set the Path property of the cookie
	Path string
	// Domain set the Domain property of the cookie
	Domain string
	// Expires set the Expires property of the cookie
	Expires time.Time
	// MaxAge set the Max-Age property of the cookie
	MaxAge int
	// Secure set the Secure property of the cookie
	Secure bool
	// SameSite set the SameSite property of the cookie
	SameSite http.SameSite
}

// InitDefaultConfig initial the default global session config
func InitDefaultConfig(c *Config) error {
	if c == nil {
		return errNilConfig
	}
	if c.Store == nil {
		return errNilStore
	}
	if len(c.Cookie.Name) == 0 {
		c.Cookie.Name = DefaultCookieName
	}
	if len(c.Cookie.Path) == 0 {
		c.Cookie.Path = "/"
	}
	if len(c.SessionPrefix) == 0 {
		c.SessionPrefix = DefaultSessionPrefix
	}
	defaultConfig = c
	return nil
}
