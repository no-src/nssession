package nssession

import (
	"errors"
	"net/http"
)

// ErrNil get nil data
var ErrNil = errors.New("nssession: nil")

// NSSession the session operation interface
type NSSession interface {
	// ID returns the session id
	ID() string

	// Get get cache data by key
	Get(k string, v any) error

	// Set set new cache data
	Set(k string, v any) error

	// Remove remove the specified key
	Remove(k string) error

	// Clear remove all the key of current session
	Clear() error
}

// Default get the session with the global session config
func Default(req *http.Request, writer http.ResponseWriter) (NSSession, error) {
	return New(defaultConfig, req, writer)
}
