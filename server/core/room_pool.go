package core

import (
	"github.com/golang/glog"
	"sync"
	"time"
)

type RoomPool struct {
	roomReadyMap  map[string]*Room
	roomGamingMap map[string]*Room
	roomEndedMap  map[string]*Room
	mu            sync.RWMutex
}

var GameRoomPool *RoomPool

func init() {
	GameRoomPool = new(RoomPool)
	GameRoomPool.roomReadyMap = make(map[string]*Room, 64)
	GameRoomPool.roomGamingMap = make(map[string]*Room, 64)
	GameRoomPool.roomEndedMap = make(map[string]*Room, 64)
}

func GetJoinableRoom() (res *Room) {
	res = GameRoomPool.getRandomReadyRoom()
	if res == nil {
		res = GameRoomPool.createRoom()
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

func (r *RoomPool) getRandomReadyRoom() *Room {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, v := range r.roomReadyMap {
		if v.IsFull() {
			continue
		}
		return v
	}
	return nil
}

func (r *RoomPool) createRoom() *Room {
	r.mu.Lock()
	defer r.mu.Unlock()
	room := generateRoom()
	r.roomReadyMap[room.RoomID] = room
	return room
}
