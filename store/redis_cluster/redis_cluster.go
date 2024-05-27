package redis_cluster

import (
	"github.com/no-src/nscache/redis_cluster"
	"github.com/no-src/nssession/store"
)

// Driver the unique name of the Redis Cluster driver
var Driver store.DriverName = redis_cluster.DriverName
