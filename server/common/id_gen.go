package common

import (
	"fmt"
	"math/rand"
	"sync/atomic"
)

var id int64
var roomId int64

func GetNextUserId() int64 {
	atomic.AddInt64(&id, 1)
	return id
}

func GetRandomRoomId() string {
	atomic.AddInt64(&roomId, 1)
	// 我愿称之为小雪花算法
	return fmt.Sprintf("R%06d%06d", roomId, rand.Int63n(1000000))
}
