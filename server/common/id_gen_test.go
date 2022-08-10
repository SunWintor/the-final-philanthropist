package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRandomRoomId(t *testing.T) {
	t.Run("测试生成的roomId是否有效", func(t *testing.T) {
		roomId := GetRandomRoomId()
		assert.Contains(t, roomId, "R0") // 笑死，这个早晚踩坑
		assert.Equal(t, len(roomId), 13)
	})
}
