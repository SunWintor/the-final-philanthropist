package core

import (
	"sync"
)

type Room struct {
	RoomID    string
	Status    int64             // 1：游戏未开始，2：游戏已开始，3：游戏已结束
	PlayerMap map[int64]*Player // list删除一个玩家简直恶心死了！我宁可每次遍历都是乱的qwq。
	mu        sync.RWMutex
}

const (
	RoomMaxPlayerCount = 8
)

const (
	GameNotStart = iota + 1
	Gaming
	GameEnded
)

func (r *Room) Join(player *Player) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.PlayerMap[player.UserId] = player
}

func (r *Room) Exit(player *Player) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.PlayerMap, player.UserId)
}

func (r *Room) GameStart() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Status = Gaming
	GameRoomPool.readyToGaming(r)
}

func (r *Room) GameEnd() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Status = GameEnded
	GameRoomPool.gamingToEnded(r)
}
