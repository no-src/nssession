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

	// DefaultSessionKey the default cookie name for session
	DefaultSessionKey ContextKey = "session-id"
	// DefaultSessionPrefix the default session prefix for session store
	DefaultSessionPrefix = "session"
)

// Config the configuration of session
type Config struct {
	// Connection the connection string of the session
	Connection string
	// Expiration the expiration time of the session
	Expiration time.Duration
	// SessionKey the cookie name for session
	SessionKey ContextKey
	// SessionPrefix the session prefix for session store
	SessionPrefix string
	// Store the session store component
	Store store.Store
}

// ContextKey the context key used for the context.WithValue
type ContextKey string

// InitDefaultConfig initial the default global session config
func InitDefaultConfig(c *Config) error {
	if c == nil {
		return errNilConfig
	}
	if c.Store == nil {
		return errNilStore
	}
	if len(c.SessionKey) == 0 {
		c.SessionKey = DefaultSessionKey
	}
	if len(c.SessionPrefix) == 0 {
		c.SessionPrefix = DefaultSessionPrefix
	}
	defaultConfig = c
	return nil
}
