# nssession

[![Build](https://img.shields.io/github/actions/workflow/status/no-src/nssession/go.yml?branch=main)](https://github.com/no-src/nssession/actions)
[![License](https://img.shields.io/github/license/no-src/nssession)](https://github.com/no-src/nssession/blob/main/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/no-src/nssession.svg)](https://pkg.go.dev/github.com/no-src/nssession)
[![Go Report Card](https://goreportcard.com/badge/github.com/no-src/nssession)](https://goreportcard.com/report/github.com/no-src/nssession)
[![codecov](https://codecov.io/gh/no-src/nssession/branch/main/graph/badge.svg?token=4KMBA2D6TY)](https://codecov.io/gh/no-src/nssession)
[![Release](https://img.shields.io/github/v/release/no-src/nssession)](https://github.com/no-src/nssession/releases)

## Installation

```bash
go get -u github.com/no-src/nssession
```

## Quick Start

```go
package main

import (
	"net/http"
	"time"

	"github.com/no-src/log"
	"github.com/no-src/nssession"
	"github.com/no-src/nssession/store"
	"github.com/no-src/nssession/store/memory"
)

func main() {
	// initial default session config
	c := &nssession.Config{
		Connection:    "memory:",
		Expiration:    time.Hour,
		CookieName:    nssession.DefaultCookieName,
		SessionPrefix: nssession.DefaultSessionPrefix,
		Store:         store.NewStore(memory.Driver),
	}
	err := nssession.InitDefaultConfig(c)
	if err != nil {
		log.Error(err, "init the default config error")
		return
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// get session component
		session, err := nssession.Default(request, writer)
		if err != nil {
			log.Error(err, "get session component error")
			return
		}

		// set session data
		k := "hello"
		var v string
		err = session.Set(k, "world")
		if err != nil {
			log.Error(err, "set session data error")
			return
		}

		// get session data
		err = session.Get(k, &v)
		if err != nil {
			log.Error(err, "get session data error")
			return
		}

		log.Info("get the session data success, k=%s v=%s", k, v)

		// remove session data
		err = session.Remove(k)
		if err != nil {
			log.Error(err, "remove session data error")
			return
		}

		// clear all session data for the current session
		err = session.Clear()
		if err != nil {
			log.Error(err, "clear session data error")
			return
		}
	})
	http.ListenAndServe(":8080", nil)
}
```