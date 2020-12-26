package internal

import (
	"time"
)

type UserInfo struct {
	Uid         uint64 `json:"uid" desc:"用户ID"`
	Ptime       int64  `json:"ptime" desc:"心跳时间"`
	LoginPlat   int    `json:"login_plat" desc:"登录平台"`
	LoginTime   int64  `json:"login_time" desc:"登录时间"`
	LastAddTime int64  `json:"last_add_time" desc:"最后添加时间"`
}

func (u *UserInfo) Clear() {
	u = &UserInfo{}
}

func (u *UserInfo) ResetLogin() {
	u.LoginTime = time.Now().Unix()
	u.LastAddTime = time.Now().Unix()
}
func (u *UserInfo) ResetHeart() {
	u.Ptime = time.Now().Unix()
}
func (u *UserInfo) FlushOnlineTime(force bool) {
	cur := time.Now().Unix()
	diff := cur - u.LastAddTime
	if diff > 30 || force {
		//save flush cache,db
		u.LastAddTime = cur
	}
}
