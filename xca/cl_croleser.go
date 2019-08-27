package xca

import (
	"cmx/xcm"
	"cmx/xsql"
	"cmx/xwb"
	"log"
	"net/http"
	"sync"
)

//NewCRoleSer 创建一个角色管理对象
func NewCRoleSer(dbstr string, dbtype string, dbcommit bool) *CRoleSer {
	return &CRoleSer{dbstr: dbstr, dbtype: dbtype, dbcommit: dbcommit}
}

//CRoleSer 角色服务对应关系
type CRoleSer struct {
	lock     sync.RWMutex
	dbstr    string
	dbtype   string
	dbcommit bool
	crss     map[string][]string
}

//query 获取服务对应的角色组
func (cr *CRoleSer) query(serid string) (roles []string, ok bool) {
	cr.lock.RLock()
	defer cr.lock.RUnlock()
	roles, ok = cr.crss[serid]
	return
}

//Query 获取服务对应的角色组
func (cr *CRoleSer) Query(r *http.Request) (roles []string, ok bool) {
	roles, ok = cr.query(xcm.GetMD5(xwb.GetURL(r) + r.FormValue("ctype")))
	return
}

//UpdateSerRole 更新角色服务对应关系
func (cr *CRoleSer) UpdateSerRole() {
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
	cr.lock.Lock()
	defer cr.lock.Unlock()
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
