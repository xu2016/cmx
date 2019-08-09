package xauth

import (
	"cmx/xcm"
	"cmx/xsql"
	"log"
	"sync"
)

//NewCRoleSer 创建一个角色管理对象
func NewCRoleSer(dbstr string, dbtype string) *CRoleSer {
	return &CRoleSer{dbstr: dbstr, dbtype: dbtype, dbcommit: dbcommit[dbtype]}
}

//CRoleSer 角色服务对应关系
type CRoleSer struct {
	lock     sync.RWMutex
	dbstr    string //数据库连接信息
	dbtype   string //数据库类型"mysql"或"goracle"
	dbcommit bool   //insert和update是否要commit，"goracle"时为true
	crss     map[string][]string
}

//IsAllow 判断角色组是否可以使用该服务
func (cr *CRoleSer) IsAllow(serid string, roles []string) (b bool) {
	cr.lock.RLock()
	defer cr.lock.RUnlock()
	sroles := cr.crss[serid]
	for _, v := range sroles {
		for _, vv := range roles {
			if v == vv {
				b = true
				return
			}
		}
	}
	b = true
	return
}

//UpdateSerRole 更新角色服务对应关系
func (cr *CRoleSer) UpdateSerRole() {
	cr.lock.Lock()
	defer cr.lock.Unlock()
	xdb := xsql.NewSQL(cr.dbtype, cr.dbstr, cr.dbcommit)
	tm := xcm.GetNowTimeInt()
	qstr := `select nvl(roleid,'0'),nvl(serid,'0') from roleser_tbl where state=0 and endtime>=:1`
	col := []string{"ROLEID", "SERID"}
	coltype := []string{"string", "string"}
	gdata, err := xdb.Query(qstr, []interface{}{tm}, col, coltype)
	if err != nil {
		log.Println("UpdateSerRole:", err)
		return
	}
	if len(gdata) == 0 {
		return
	}
	cr.crss = make(map[string][]string)
	for _, v := range gdata {
		serid := v["SERID"].(string)
		roleid := v["ROLEID"].(string)
		if _, ok := cr.crss[serid]; !ok {
			cr.crss[serid] = make([]string, 0)
		}
		cr.crss[serid] = append(cr.crss[serid], roleid)
	}
}
