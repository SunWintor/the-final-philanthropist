package room

import (
	"github.com/SunWintor/tfp/server/common"
	"github.com/SunWintor/tfp/server/core"
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/pkg/errors"
	"sync"
)

type Room struct {
	RoomId    string
	Status    int64                  // 1：游戏未开始，2：游戏已开始，3：游戏已结束
	playerMap map[int64]*core.Player // list删除一个玩家简直恶心死了！我宁可每次遍历都是乱的qwq。不导出是为了外面不要随便修改，可能导致panic
	mu        sync.RWMutex           // 为什么一个锁？问就是懒233
}

const (
	RoomMaxPlayerCount = 8
)

const (
	GameReady = int64(iota) + 1
	Gaming
	GameEnded
)

// key userId
// value roomId
var userCurrentRoomIdMap sync.Map

func init() {
	userCurrentRoomIdMap = sync.Map{}
}

func generateReadyRoom() *Room {
	room := new(Room)
	room.RoomId = common.GetRandomRoomId()
	room.Status = GameReady
	room.playerMap = make(map[int64]*core.Player, RoomMaxPlayerCount)
	return room
}

func GetCurrentUserRoomId(userId int64) string {
	roomId, _ := userCurrentRoomIdMap.Load(userId)
	if roomId == nil {
		return ""
	}
	return roomId.(string)
}

func (r *Room) IsFull() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.playerMap) == RoomMaxPlayerCount
}

func (r *Room) Join(player *core.Player) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.IsFull() {
		return errors.WithMessagef(ecode.RoomIsFullError, "player:%v, room:%v", player, r)
	}
	if player.RoomId != "" {
		return errors.WithMessagef(ecode.AlreadyInRoomError, "player:%v, room:%v", player, r)
	}
	r.playerMap[player.UserId] = player
	player.RoomId = r.RoomId
	userCurrentRoomIdMap.Store(player.UserId, r.RoomId)
	return nil
}

func (r *Room) Exit(player *core.Player) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if player.RoomId != r.RoomId {
		return errors.WithMessagef(ecode.PlayerNotInRoomError, "player:%v, room:%v", player, r)
	}
	delete(r.playerMap, player.UserId)
	player.RoomId = ""
	userCurrentRoomIdMap.Delete(player.UserId)
	return nil
}

func (r *Room) GameStart() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Status = Gaming
	gameRoomPool.readyToGaming(r)
}

func (r *Room) GameEnd() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Status = GameEnded
	userCurrentRoomIdMap.Range(func(key, value interface{}) bool {
		if value.(string) == r.RoomId {
			// sync.Map的遍历删除是完备的，不必担心。
			userCurrentRoomIdMap.Delete(key)
		}
		return true
	})
	gameRoomPool.gamingToEnded(r)
}
