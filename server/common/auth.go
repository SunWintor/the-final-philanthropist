package common

import "sync"

type LoginUserCache struct {
	loginUser map[string]string
	mu        sync.RWMutex
}

var LoginUser *LoginUserCache

// 因为项目早期并不想要持久化，所以用户的登录完全由内存维护
// 未来需要持久化的时候，这里会进行变更，使用其他方案来实现
func init() {
	LoginUser = new(LoginUserCache)
	LoginUser.loginUser = make(map[string]string, 64)
}

func (l *LoginUserCache) Put(username, token string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.loginUser[username] = token
}

func (l *LoginUserCache) Get(username string) (token string, ok bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	token, ok = l.loginUser[username]
	return
}

func (l *LoginUserCache) Delete(username string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.loginUser, username)
}
