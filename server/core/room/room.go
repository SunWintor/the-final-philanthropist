package room

import (
	"github.com/SunWintor/tfp/server/common"
	"github.com/SunWintor/tfp/server/core/game"
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/SunWintor/tfp/server/model"
	"github.com/pkg/errors"
	"sync"
)

type Room struct {
	RoomId  string
	Status  int64 // 1：游戏未开始，2：游戏已开始，3：游戏已结束
	Game    *game.Game
	userMap map[int64]*RoomUser // list删除一个玩家简直恶心死了！我宁可每次遍历都是乱的qwq。不导出是为了外面不要随便修改，可能导致panic
	mu      sync.RWMutex        // 为什么一个锁？问就是懒233
}

const (
	RoomUserLimit = 8
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
	room.userMap = make(map[int64]*RoomUser, RoomUserLimit)
	return room
}

func GetUserRoomId(userId int64) string {
	roomId, _ := userCurrentRoomIdMap.Load(userId)
	if roomId == nil {
		return ""
	}
	return roomId.(string)
}

func UserRoom(userId int64) (r *Room, err error) {
	roomId := GetUserRoomId(userId)
	if roomId == "" {
		err = ecode.PlayerNotInRoomError
		return
	}
	r = GetRoom(roomId)
	return
}

func (r *Room) Ready(userId int64, ready bool) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	var roomUser *RoomUser
	var ok bool
	if roomUser, ok = r.userMap[userId]; !ok {
		return ecode.PlayerNotInRoomError
	}
	roomUser.IsReady = ready
	return nil
}

func (r *Room) isFull() bool {
	return len(r.userMap) == RoomUserLimit
}

func (r *Room) Join(roomUser *RoomUser) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.isFull() {
		return errors.WithMessagef(ecode.RoomIsFullError, "roomUser:%v, room:%v", roomUser, r)
	}
	if roomUser.RoomId != "" {
		return errors.WithMessagef(ecode.AlreadyInRoomError, "roomUser:%v, room:%v", roomUser, r)
	}
	r.userMap[roomUser.UserId] = roomUser
	roomUser.RoomId = r.RoomId
	userCurrentRoomIdMap.Store(roomUser.UserId, r.RoomId)
	return nil
}

func (r *Room) Exit(userId int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	roomUser, ok := r.userMap[userId]
	if !ok {
		return ecode.PlayerNotInRoomError
	}
	delete(r.userMap, roomUser.UserId)
	roomUser.RoomId = ""
	userCurrentRoomIdMap.Delete(roomUser.UserId)
	return nil
}

func (r *Room) GameStart() (endChan <-chan struct{}, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if len(r.userMap) < 2 {
		return nil, ecode.InsufficientPlayerError
	}
	for _, roomUser := range r.userMap {
		if !roomUser.IsReady {
			return nil, ecode.PlayerNotReadyError
		}
	}
	endChan = r.gameInit()
	r.Status = Gaming
	gameRoomPool.readyToGaming(r)
	return endChan, nil
}

func (r *Room) gameInit() <-chan struct{} {
	r.Game = &game.Game{
		GameId: common.GetRandomGameId(),
		RoomId: r.RoomId,
	}
	playerMap := make(map[string]*game.Player, len(r.userMap)*2)
	for _, value := range r.userMap {
		player := value.ToPlayer()
		playerMap[player.PlayerId] = player
	}
	r.Game.GameInit(playerMap)
	endChan := r.Game.Start()
	go func() {
		<-endChan
		r.gameEnd()
	}()
	return endChan
}

func (r *Room) gameEnd() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Status = GameEnded
	gameRoomPool.gamingToEnded(r)
}

func (r *Room) ToReply() *model.RoomInfoReply {
	r.mu.RLock()
	defer r.mu.RUnlock()
	gameId := ""
	if r.Game != nil {
		gameId = r.Game.GameId
	}
	res := &model.RoomInfoReply{
		RoomId: r.RoomId,
		GameId: gameId,
		Status: r.Status,
	}
	for _, roomUser := range r.userMap {
		res.RoomUserList = append(res.RoomUserList, roomUser.ToReply())
	}
	return res
}
