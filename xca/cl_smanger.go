package xca

import (
	"cmx/xcm"
	"errors"
	"sync"
	"time"
)

//SManager Session管理器
type SManager struct {
	lock        sync.RWMutex       // protects session
	sid         map[string]session //session id 唯一标示
	maxlifetime int64
}

//session session存储结构
type session struct {
	uid          string    //用户账号
	timeAccessed time.Time //最后访问时间
	rid          []string  //角色组
	phone        string    //用户手机号
	city         string    //地市
}

//NewSManager 参加一个Session管理器
func NewSManager(maxlifetime int64) *SManager {
	sid := make(map[string]session, 0)
	return &SManager{sid: sid, maxlifetime: maxlifetime}
}

//GC 定时对过期的Session进行删除
func (sm *SManager) GC() {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	for k, v := range sm.sid {
		if (v.timeAccessed.Unix() + sm.maxlifetime) < time.Now().Unix() {
			delete(sm.sid, k)
		}
	}
	time.AfterFunc(time.Duration(sm.maxlifetime*2), func() { sm.GC() })
}

//UserIsLogin 判断用户是否登陆
func (sm *SManager) UserIsLogin(sid string) bool {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	zs, ok := sm.sid[sid]
	if !ok {
		return false
	}
	if (zs.timeAccessed.Unix() + sm.maxlifetime) < time.Now().Unix() {
		return false
	}
	zs.timeAccessed = time.Now()
	sm.sid[sid] = zs
	return true
}

//AddSession 用户登陆添加Session
func (sm *SManager) AddSession(userid, phone, city string, rid []string) (sid string, err error) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	sid = xcm.GetMD5(phone + userid + city + xcm.GetRandomString(8, xcm.NSDSTR))
	if sid == "" {
		err = errors.New("Add  session error")
		return
	}
	zs := session{uid: userid, phone: phone, rid: rid, city: city, timeAccessed: time.Now()}
	sm.sid[sid] = zs
	return
}

//DelSession 用户注销删除Session
func (sm *SManager) DelSession(sid string) (err error) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	delete(sm.sid, sid)
	return
}

//GetUserID 获取用户ID
func (sm *SManager) GetUserID(sid string) (uid string, err error) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	zs, ok := sm.sid[sid]
	if !ok {
		err = errors.New("No session")
		return
	}
	uid = zs.uid
	return
}

//GetUserPhone 获取用户phone
func (sm *SManager) GetUserPhone(sid string) (phone string, err error) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	zs, ok := sm.sid[sid]
	if !ok {
		err = errors.New("No session")
		return
	}
	phone = zs.phone
	return
}

//GetUserCity 获取用户所属地市
func (sm *SManager) GetUserCity(sid string) (city string) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	zs, ok := sm.sid[sid]
	if !ok {
		return
	}
	city = zs.city
	return
}

//GetUserRoles 获取用户角色组
func (sm *SManager) GetUserRoles(sid string) (rid []string) {
	sm.lock.RLock()
	defer sm.lock.RUnlock()
	zs, ok := sm.sid[sid]
	if !ok {
		return
	}
	rid = zs.rid
	return
}
