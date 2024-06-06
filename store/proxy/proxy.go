package proxy

import (
	"github.com/no-src/nscache/proxy/client"
	"github.com/no-src/nssession/store"
)

// Driver the unique name of the proxy driver
var Driver store.DriverName = client.DriverName
