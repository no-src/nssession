package etcd

import (
	"github.com/no-src/nscache/etcd"
	"github.com/no-src/nssession/store"
)

// Driver the unique name of the Etcd driver
var Driver store.DriverName = etcd.DriverName
