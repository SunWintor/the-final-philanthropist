package core

import (
	"github.com/SunWintor/tfp/server/common"
	"github.com/golang/glog"
	"sync"
	"time"
)

type RoomPool struct {
	roomReadyMap  map[string]*Room
	roomGamingMap map[string]*Room
	roomEndedMap  map[string]*Room
	mu            sync.RWMutex // todo 现在共享一把锁纯粹因为懒，房间操作不是频繁操作，所以目前呆胶布
}

var GameRoomPool *RoomPool

func init() {
	GameRoomPool = new(RoomPool)
	GameRoomPool.roomReadyMap = make(map[string]*Room, 64)
	GameRoomPool.roomGamingMap = make(map[string]*Room, 64)
	GameRoomPool.roomEndedMap = make(map[string]*Room, 64)
}

func (r *RoomPool) GetNotFullRoom() (res *Room) {
	res = r.getRandomFreeRoom()
	if res == nil {
		res = r.createRoom()
	}
	return
}

func (r *RoomPool) readyToGaming(room *Room) (ok bool) {
	if room == nil {
		glog.Error("readyToGaming room is nil ! %v", r)
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok = r.roomReadyMap[room.RoomID]; ok {
		r.roomGamingMap[room.RoomID] = room
		delete(r.roomReadyMap, room.RoomID)
	}
	return
}

func (r *RoomPool) gamingToEnded(room *Room) (ok bool) {
	if room == nil {
		glog.Error("gamingToEnded room is nil ! %v", r)
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok = r.roomGamingMap[room.RoomID]; ok {
		r.roomEndedMap[room.RoomID] = room
		delete(r.roomGamingMap, room.RoomID)
		go func() {
			<-time.After(time.Minute * 10)
			delete(r.roomEndedMap, room.RoomID)
		}()
	}
	return
}

func (r *RoomPool) getRandomFreeRoom() *Room {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, v := range r.roomReadyMap {
		if len(v.PlayerMap) == RoomMaxPlayerCount {
			continue
		}
		return v
	}
	return nil
}

func (r *RoomPool) createRoom() *Room {
	r.mu.Lock()
	defer r.mu.Unlock()
	room := new(Room)
	room.RoomID = common.GetRandomRoomId()
	room.Status = GameNotStart
	room.PlayerMap = make(map[int64]*Player, RoomMaxPlayerCount)
	return room
}
