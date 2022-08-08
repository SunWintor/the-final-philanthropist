package common

import (
	"github.com/SunWintor/tfp/server/model"
	"sync"
)

type UserCache struct {
	nameMap map[string]*model.UserInfo
	idMap   map[int64]*model.UserInfo
	mu      sync.RWMutex
}

var User *UserCache

// 因为项目早期并不想要持久化，所以用户的登录完全由内存维护
// 未来需要持久化的时候，这里会进行变更，使用其他方案来实现
func init() {
	User = new(UserCache)
	User.nameMap = make(map[string]*model.UserInfo, 64)
	User.idMap = make(map[int64]*model.UserInfo, 64)
}

func (u *UserCache) Put(userInfo *model.UserInfo) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.nameMap[userInfo.Username] = userInfo
	u.idMap[userInfo.Id] = userInfo
}

// GetCopyByName 获取的是副本，要改用put。
func (u *UserCache) GetCopyByName(username string) (userInfo *model.UserInfo, ok bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	var tmp *model.UserInfo
	tmp, ok = u.nameMap[username]
	userInfo = tmp.DeepCopy()
	return
}

// GetCopyById 获取的是副本，要改用put
func (u *UserCache) GetCopyById(id int64) (userInfo *model.UserInfo, ok bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	var tmp *model.UserInfo
	userInfo, ok = u.idMap[id]
	userInfo = tmp.DeepCopy()
	return
}

func (u *UserCache) Delete(username string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	delete(u.nameMap, username)
}
