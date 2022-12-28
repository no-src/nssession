package store

import (
	"errors"
	"strings"
	"sync"

	"github.com/no-src/nscache"
)

var (
	errInvalidStoreDriver     = errors.New("invalid store driver")
	errUnsupportedStoreDriver = errors.New("unsupported store driver")
)

// Store the session store component
type Store interface {
	// NewCache create an instance of the store component by the specified connection string
	NewCache(conn string) (nscache.NSCache, error)
}

// DriverName the unique name of the specified driver
type DriverName string

type store struct {
	drivers []DriverName
	caches  map[string]nscache.NSCache
	mu      sync.RWMutex
}

func (s *store) NewCache(conn string) (nscache.NSCache, error) {
	s.mu.RLock()
	cache := s.caches[conn]
	s.mu.RUnlock()
	if cache != nil {
		return cache, nil
	}
	args := strings.Split(conn, ":")
	if len(args) < 2 {
		return nil, errInvalidStoreDriver
	}
	supported := false
	for _, driver := range s.drivers {
		if strings.ToLower(args[0]) == strings.ToLower(string(driver)) {
			supported = true
		}
	}
	if !supported {
		return nil, errUnsupportedStoreDriver
	}
	cache, err := nscache.NewCache(conn)
	if err != nil {
		return nil, err
	}
	s.mu.Lock()
	s.caches[conn] = cache
	s.mu.Unlock()
	return cache, nil
}

// NewStore create an instance of the Store with the specified drivers
func NewStore(drivers ...DriverName) Store {
	return &store{
		drivers: drivers,
		caches:  make(map[string]nscache.NSCache),
	}
}
