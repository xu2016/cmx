package xca

import (
	"cmx/xcm"
	"sync"
	"time"
)

//CyzmCache 验证码缓存管理器
type CyzmCache struct {
	lock        sync.Mutex           // protects session
	yzm         map[string]time.Time //session id 唯一标示
	maxlifetime int64
}

//NewCyzmCache 参加一个验证码管理器
func NewCyzmCache(maxlifetime int64) *CyzmCache {
	yzm := make(map[string]time.Time, 0)
	return &CyzmCache{yzm: yzm, maxlifetime: maxlifetime}
}

//GC 定时对过期的验证码进行删除
func (yzmm *CyzmCache) GC() {
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
func (yzmm *CyzmCache) Add(yzm, uid string) {
	yzmm.lock.Lock()
	defer yzmm.lock.Unlock()
	zsstr := xcm.GetMD5(yzm + uid)
	yzmm.yzm[zsstr] = time.Now()
	return
}

//Query 验证验证码是否正确
func (yzmm *CyzmCache) Query(yzm, uid string) bool {
	yzmm.lock.Lock()
	defer yzmm.lock.Unlock()
	yzmid := xcm.GetMD5(yzm + uid)
	if yzmTime, ok := yzmm.yzm[yzmid]; ok {
		if (yzmTime.Unix() + 300) < time.Now().Unix() {
			delete(yzmm.yzm, yzmid)
			return false
		}
		return true
	}
	return false
}

//Del 删除验证码
func (yzmm *CyzmCache) Del(yzm, uid string) {
	yzmm.lock.Lock()
	defer yzmm.lock.Unlock()
	yzmid := xcm.GetMD5(yzm + uid)
	if _, ok := yzmm.yzm[yzmid]; ok {
		delete(yzmm.yzm, yzmid)
	}
}
