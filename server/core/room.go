package core

import (
	"github.com/SunWintor/tfp/server/common"
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/pkg/errors"
	"sync"
)

type Room struct {
	RoomID    string
	Status    int64             // 1：游戏未开始，2：游戏已开始，3：游戏已结束
	playerMap map[int64]*Player // list删除一个玩家简直恶心死了！我宁可每次遍历都是乱的qwq。不导出是为了外面不要随便修改，可能导致panic
	mu        sync.RWMutex      // 为什么一个锁？问就是懒233
}

const (
	RoomMaxPlayerCount = 8
)

const (
	GameReady = iota + 1
	Gaming
	GameEnded
)

var userIdRoomIdMap sync.Map

func init() {
	userIdRoomIdMap = sync.Map{}
}

func generateRoom() *Room {
	room := new(Room)
	room.RoomID = common.GetRandomRoomId()
	room.Status = GameReady
	room.playerMap = make(map[int64]*Player, RoomMaxPlayerCount)
	return room
}

func GetRoomIdByUserId(userId int64) string {
	roomId, _ := userIdRoomIdMap.Load(userId)
	return roomId.(string)
}

func (r *Room) IsFull() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.playerMap) == RoomMaxPlayerCount
}

func (r *Room) Join(player *Player) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if player.RoomID != "" {
		return errors.WithMessagef(ecode.AlreadyInRoomError, "player:%v, room:%v", player, r)
	}
	r.playerMap[player.UserID] = player
	player.RoomID = r.RoomID
	userIdRoomIdMap.Store(player.UserID, r.RoomID)
	return nil
}

func (r *Room) Exit(player *Player) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if player.RoomID != r.RoomID {
		return errors.WithMessagef(ecode.PlayerNotInRoomError, "player:%v, room:%v", player, r)
	}
	delete(r.playerMap, player.UserID)
	player.RoomID = ""
	userIdRoomIdMap.Delete(player.UserID)
	return nil
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
	userIdRoomIdMap.Range(func(key, value interface{}) bool {
		if value.(string) == r.RoomID {
			// sync.Map的遍历删除是完备的，不必担心。
			userIdRoomIdMap.Delete(key)
		}
		return true
	})
	GameRoomPool.gamingToEnded(r)
}
