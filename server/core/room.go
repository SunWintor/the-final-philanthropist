package core

import (
	"sync"
)

type Room struct {
	RoomID    string
	Status    int64 // 1：游戏未开始，2：游戏已开始，3：游戏已结束
	PlayerMap map[int64]*Player
	mu        sync.RWMutex
}

// 维护所有的房间
var RoomMap map[string]*Room

func init() {
	RoomMap = make(map[string]*Room, 64)
}
