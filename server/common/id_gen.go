package common

import "sync/atomic"

var id int64

func GetIdentifier() int64 {
	atomic.AddInt64(&id, 1)
	return id
}
