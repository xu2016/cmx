package xauth

import (
	"cmx/xcm"
	"sync"
	"time"
)

//Cyzm 验证码管理器
type Cyzm struct {
	lock        sync.Mutex           // protects session
	yzm         map[string]time.Time //session id 唯一标示
	maxlifetime int64
}

//NewCyzm 参加一个验证码管理器
func NewCyzm(maxlifetime int64) *Cyzm {
	yzm := make(map[string]time.Time, 0)
	return &Cyzm{yzm: yzm, maxlifetime: maxlifetime}
}

//GC 定时对过期的验证码进行删除
func (yzmm *Cyzm) GC() {
	yzmm.lock.Lock()
	defer yzmm.lock.Unlock()
	for k, v := range yzmm.yzm {
		if (v.Unix() + yzmm.maxlifetime) < time.Now().Unix() {
			delete(yzmm.yzm, k)
		}
	}
	time.AfterFunc(time.Duration(yzmm.maxlifetime*2), func() { yzmm.GC() })
}

//Add 添加验证码md5(yzm+uid)
func (yzmm *Cyzm) Add(yzm, uid string) {
	yzmm.lock.Lock()
	defer yzmm.lock.Unlock()
	zsstr := xcm.GetMD5(yzm + uid)
	yzmm.yzm[zsstr] = time.Now()
	return
}

//Query 验证验证码是否正确
func (yzmm *Cyzm) Query(yzm, uid string) bool {
	yzmm.lock.Lock()
	defer yzmm.lock.Unlock()
	yzmid := xcm.GetMD5(yzm + uid)
	if yzmTime, ok := yzmm.yzm[yzmid]; ok {
		if (yzmTime.Unix() + 600) < time.Now().Unix() {
			delete(yzmm.yzm, yzmid)
			return false
		}
		return true
	}
	return false
}
