package fastcache

import (
	"github.com/no-src/nscache/fastcache"
	"github.com/no-src/nssession/store"
)

// Driver the unique name of the fastcache driver
var Driver store.DriverName = fastcache.DriverName
