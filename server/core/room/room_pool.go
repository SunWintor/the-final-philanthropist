package room

import (
	"fmt"
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
	gameRoomPool.roomEndedMap = make(map[string]*Room, 64)
}

func GetJoinableRoom() (res *Room) {
	res = pool().getRandomReadyRoom()
	if res == nil {
		res = pool().createEmptyRoomToReady()
	}
	return
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
	if room, ok = r.roomEndedMap[roomId]; ok {
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

func (r *RoomPool) gamingToEnded(room *Room) (ok bool) {
	if room == nil {
		glog.Error(fmt.Sprintf("gamingToEnded room is nil ! %v", r))
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok = r.roomGamingMap[room.RoomId]; ok {
		r.roomEndedMap[room.RoomId] = room
		delete(r.roomGamingMap, room.RoomId)
		go func() {
			<-time.After(time.Minute * 10)
			delete(r.roomEndedMap, room.RoomId)
		}()
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
