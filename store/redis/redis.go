package redis

import (
	"github.com/no-src/nscache/redis"
	"github.com/no-src/nssession/store"
)

// Driver the unique name of the Redis driver
var Driver store.DriverName = redis.DriverName
