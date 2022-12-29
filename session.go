package nssession

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/no-src/nscache"
)

type session struct {
	c      *Config
	cache  nscache.NSCache
	id     string
	req    *http.Request
	writer http.ResponseWriter
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
func New(c *Config, req *http.Request, writer http.ResponseWriter) (NSSession, error) {
	if c == nil {
		return nil, errNilConfig
	}
	s := &session{
		c:      c,
		req:    req,
		writer: writer,
	}

	var sessionID string
	if s.req != nil {
		cookie, err := s.req.Cookie(c.CookieName)
		if err == nil && cookie != nil {
			sessionID = cookie.Value
		}
	}

	if len(sessionID) == 0 {
		sessionID = s.generateID()
		if s.writer != nil {
			http.SetCookie(s.writer, &http.Cookie{
				Name:     c.CookieName,
				Value:    sessionID,
				Path:     "/",
				HttpOnly: true,
			})
		}
	}

	s.id = sessionID
	var err error
	s.cache, err = c.Store.NewCache(c.Connection)
	if err != nil {
		return nil, err
	}
	return s, nil
}
