package boltdb

import (
	"github.com/no-src/nscache/boltdb"
	"github.com/no-src/nssession/store"
)

// Driver the unique name of the BoltDB driver
var Driver store.DriverName = boltdb.DriverName
