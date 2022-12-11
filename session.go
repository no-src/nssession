package nssession

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/no-src/nscache"
)

type session struct {
	c     *Config
	cache nscache.NSCache
	id    string
}

func (s *session) ID() string {
	return s.id
}

func (s *session) Get(k string, v any) error {
	sk := s.key()
	var sd sessionData
	err := s.cache.Get(sk, &sd)
	if err != nil {
		if err == nscache.ErrNil {
			err = ErrNil
		}
		return err
	}
	d := sd.Data[k]
	if d == nil {
		return ErrNil
	}
	vBytes, err := json.Marshal(d)
	if err == nil {
		err = json.Unmarshal(vBytes, &v)
	}
	return err
}

func (s *session) Set(k string, v any) error {
	sk := s.key()
	var sd sessionData
	err := s.cache.Get(sk, &sd)
	if err == nscache.ErrNil {
		sd = sessionData{Data: make(map[string]any)}
		err = nil
	}
	if err == nil {
		sd.Data[k] = v
		err = s.cache.Set(sk, sd, s.c.Expiration)
	}
	return err
}

func (s *session) Remove(k string) error {
	sk := s.key()
	var sd sessionData
	err := s.cache.Get(sk, &sd)
	if err != nil {
		if err == nscache.ErrNil {
			err = nil
		}
		return err
	}
	delete(sd.Data, k)
	return s.cache.Set(sk, sd, s.c.Expiration)
}

func (s *session) Clear() error {
	sk := s.key()
	return s.cache.Remove(sk)
}

func (s *session) generateID() string {
	return uuid.NewString()
}

func (s *session) key() string {
	return fmt.Sprintf("%s_%s", s.c.SessionPrefix, s.id)
}

// New get the session with the specified session config
func New(ctx context.Context, c *Config) (NSSession, error) {
	if c == nil {
		return nil, errNilConfig
	}
	s := &session{
		c: c,
	}
	sessionID := ctx.Value(c.SessionKey)
	if sessionID == nil {
		s.id = s.generateID()
	} else {
		s.id = sessionID.(string)
		ctx = context.WithValue(ctx, c.SessionKey, s.id)
	}
	var err error
	s.cache, err = c.Store.NewCache(c.Connection)
	if err != nil {
		return nil, err
	}
	return s, nil
}
