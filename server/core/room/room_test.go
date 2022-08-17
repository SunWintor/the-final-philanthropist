package room

import (
	"testing"
)

func TestGetCurrentUserRoomId(t *testing.T) {
	type args struct {
		userId int64
	}
	existsUser := int64(2)
	existsRoomId := "test"
	userCurrentRoomIdMap.Store(existsUser, existsRoomId)
	tests := []struct {
		name string
		args args
		want string
	}{
		{"获取一个不存在的user", args{1}, ""},
		{"获取一个存在的user", args{existsUser}, existsRoomId},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUserRoomId(tt.args.userId); got != tt.want {
				t.Errorf("GetUserRoomId() = %v, want %v", got, tt.want)
			}
		})
	}
}
