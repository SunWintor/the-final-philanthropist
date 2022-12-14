package common

import (
	"fmt"
	"math/rand"
	"sync/atomic"
)

var userId int64
var roomId int64
var playerId int64
var gameId int64

func GetNextUserId() int64 {
	atomic.AddInt64(&userId, 1)
	return userId
}

func GetRoomId() string {
	return id("R", &roomId)
}

func GetRandomPlayerId() string {
	return randomId("P", &playerId)
}

func GetRandomGameId() string {
	return randomId("G", &gameId)
}

func randomId(prefix string, id *int64) string {
	atomic.AddInt64(id, 1)
	// 我愿称之为小雪花算法
	return fmt.Sprintf("%s%06d%06d", prefix, *id, rand.Int63n(1000000))
}

func id(prefix string, id *int64) string {
	atomic.AddInt64(id, 1)
	return fmt.Sprintf("%s%06d", prefix, *id)
}
