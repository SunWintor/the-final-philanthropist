package room

import (
	"fmt"
	"github.com/golang/glog"
	"sync"
)

type RoomPool struct {
	roomReadyMap  map[string]*Room
	roomGamingMap map[string]*Room
	mu            sync.RWMutex
}

var gameRoomPool *RoomPool
var once sync.Once

func pool() *RoomPool {
	if gameRoomPool == nil {
		once.Do(func() {
			roomPoolInit()
		})
	}
	return gameRoomPool
}

func roomPoolInit() {
	gameRoomPool = new(RoomPool)
	gameRoomPool.roomReadyMap = make(map[string]*Room, 64)
	gameRoomPool.roomGamingMap = make(map[string]*Room, 64)
}

func GetJoinableRoom() (res *Room) {
	res = pool().getRandomReadyRoom()
	if res == nil {
		res = pool().createEmptyRoomToReady()
	}
	return
}

func CreateEmptyRoom() (res *Room) {
	return pool().createEmptyRoomToReady()
}

func GetRoom(roomId string) (room *Room) {
	r := pool()
	r.mu.RLock()
	defer r.mu.RUnlock()
	var ok bool
	if room, ok = r.roomReadyMap[roomId]; ok {
		return
	}
	if room, ok = r.roomGamingMap[roomId]; ok {
		return
	}
	return
}

func (r *RoomPool) readyToGaming(room *Room) (ok bool) {
	if room == nil {
		glog.Error(fmt.Sprintf("readyToGaming room is nil ! %v", r))
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok = r.roomReadyMap[room.RoomId]; ok {
		r.roomGamingMap[room.RoomId] = room
		delete(r.roomReadyMap, room.RoomId)
	}
	return
}

func (r *RoomPool) gamingToReady(room *Room) (ok bool) {
	if room == nil {
		glog.Error(fmt.Sprintf("readyToGaming room is nil ! %v", r))
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok = r.roomGamingMap[room.RoomId]; ok {
		r.roomReadyMap[room.RoomId] = room
		delete(r.roomGamingMap, room.RoomId)
	}
	return
}

func (r *RoomPool) readyToEnded(room *Room) (ok bool) {
	if room == nil {
		glog.Error(fmt.Sprintf("readyToEnded room is nil ! %v", r))
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok = r.roomReadyMap[room.RoomId]; ok {
		delete(r.roomReadyMap, room.RoomId)
	}
	return
}

func (r *RoomPool) getRandomReadyRoom() *Room {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, v := range r.roomReadyMap {
		if v.isFull() {
			continue
		}
		return v
	}
	return nil
}

func (r *RoomPool) createEmptyRoomToReady() *Room {
	r.mu.Lock()
	defer r.mu.Unlock()
	room := generateReadyRoom()
	r.roomReadyMap[room.RoomId] = room
	return room
}
