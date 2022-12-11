package buntdb

import (
	"github.com/no-src/nscache/buntdb"
	"github.com/no-src/nssession/store"
)

// Driver the unique name of the BuntDB driver
var Driver store.DriverName = buntdb.DriverName
