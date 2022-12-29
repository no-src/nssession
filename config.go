package nssession

import (
	"errors"
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
	// CookieName the cookie name for session
	CookieName string
	// SessionPrefix the session prefix for session store
	SessionPrefix string
	// Store the session store component
	Store store.Store
}

// InitDefaultConfig initial the default global session config
func InitDefaultConfig(c *Config) error {
	if c == nil {
		return errNilConfig
	}
	if c.Store == nil {
		return errNilStore
	}
	if len(c.CookieName) == 0 {
		c.CookieName = DefaultCookieName
	}
	if len(c.SessionPrefix) == 0 {
		c.SessionPrefix = DefaultSessionPrefix
	}
	defaultConfig = c
	return nil
}
