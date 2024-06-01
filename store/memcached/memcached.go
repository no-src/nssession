package memcached

import (
	"github.com/no-src/nscache/memcached"
	"github.com/no-src/nssession/store"
)

// Driver the unique name of the memcached driver
var Driver store.DriverName = memcached.DriverName
