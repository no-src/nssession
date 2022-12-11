package memory

import (
	"github.com/no-src/nscache/memory"
	"github.com/no-src/nssession/store"
)

// Driver the unique name of the memory driver
var Driver store.DriverName = memory.DriverName
