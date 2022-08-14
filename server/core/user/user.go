package user

import (
	"sync"

	"github.com/SunWintor/tfp/server/model"
)

type UserCache struct {
	nameMap map[string]*model.UserInfo
	IdMap   map[int64]*model.UserInfo
	mu      sync.RWMutex
}

var user *UserCache

// 因为项目早期并不想要持久化，所以用户的登录完全由内存维护
// 未来需要持久化的时候，这里会进行变更，使用其他方案来实现
func init() {
	user = new(UserCache)
	user.nameMap = make(map[string]*model.UserInfo, 64)
	user.IdMap = make(map[int64]*model.UserInfo, 64)
}

func Put(userInfo *model.UserInfo) {
	user.mu.Lock()
	defer user.mu.Unlock()
	user.nameMap[userInfo.Username] = userInfo
	user.IdMap[userInfo.Id] = userInfo
}

// GetCopyByName 获取的是副本，如果想要修改获取到的数据，需要使用put。
func GetCopyByName(username string) (userInfo *model.UserInfo, ok bool) {
	user.mu.RLock()
	defer user.mu.RUnlock()
	var tmp *model.UserInfo
	tmp, ok = user.nameMap[username]
	userInfo = tmp.DeepCopy()
	return
}

// GetCopyById 获取的是副本，如果想要修改获取到的数据，需要使用put。
func GetCopyById(Id int64) (userInfo *model.UserInfo, ok bool) {
	user.mu.RLock()
	defer user.mu.RUnlock()
	var tmp *model.UserInfo
	userInfo, ok = user.IdMap[Id]
	userInfo = tmp.DeepCopy()
	return
}

func Delete(username string) {
	user.mu.Lock()
	defer user.mu.Unlock()
	delete(user.nameMap, username)
}
