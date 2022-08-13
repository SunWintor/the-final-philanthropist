package room

import (
	"github.com/SunWintor/tfp/server/core"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGameRoomPool(t *testing.T) {
	t.Run("初始化roomPool", func(t *testing.T) {
		roomPoll := GameRoomPool()
		assert.NotNil(t, roomPoll, gameRoomPool)
	})
}

func TestGetJoinableRoom(t *testing.T) {
	emptyRoomId := "emptyRoomid"
	empty := &Room{RoomId: emptyRoomId, playerMap: map[int64]*core.Player{}}

	fullRoomId := "fullRoomid"
	full := &Room{RoomId: fullRoomId, playerMap: map[int64]*core.Player{}}
	for i := int64(0); i < RoomMaxPlayerCount; i++ {
		full.playerMap[i] = &core.Player{PlayerId: i}
	}
	full.playerMap[1] = &core.Player{}

	t.Run("没有一个有效的房间，生成新房间。", func(t *testing.T) {
		roomPoolInit()
		got := GetJoinableRoom()
		assert.Equal(t, len(got.playerMap), 0)
		assert.Equal(t, got.Status, GameReady)
	})
	t.Run("存在有效的房间，返回有效房间。", func(t *testing.T) {
		roomPoolInit()
		gameRoomPool.roomReadyMap[emptyRoomId] = empty
		got := GetJoinableRoom()
		if !reflect.DeepEqual(got, empty) {
			t.Errorf("GetJoinableRoom() = %v, want %v", got, empty)
		}
	})
	t.Run("存在满员有效房间，返回新房间。", func(t *testing.T) {
		roomPoolInit()
		gameRoomPool.roomReadyMap[fullRoomId] = full
		got := GetJoinableRoom()
		assert.NotEqual(t, got.RoomId, fullRoomId)
		assert.Equal(t, len(got.playerMap), 0)
		assert.Equal(t, got.Status, GameReady)
	})
	t.Run("存在满员有效房间和一个空闲房间，返回空闲房间。", func(t *testing.T) {
		roomPoolInit()
		gameRoomPool.roomReadyMap[emptyRoomId] = empty
		gameRoomPool.roomReadyMap[fullRoomId] = full
		got := GetJoinableRoom()
		assert.NotEqual(t, got.RoomId, fullRoomId)
		if !reflect.DeepEqual(got, empty) {
			t.Errorf("GetJoinableRoom() = %v, want %v", got, empty)
		}
	})
}
