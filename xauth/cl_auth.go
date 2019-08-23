package xauth

import (
	"cmx/xca"
	"cmx/xtx"
	"net/http"
)

//CAuth 权限设置类
type CAuth struct {
	pcCName string
	wxCName string
	h5CName string
	dbstr   string
	dbtype  string
	Gsm     *xca.SManager  //Gpcsm 全局PC端Session管理器
	Gyzm    *xca.CyzmCache //Gyzm 全局验证码管理器
	Guk     *xca.CUniKey   //Guk 全局唯一key管理器
	Gcrs    *xca.CRoleSer  //Gcrs 全局角色服务对应关系管理对象
	Gctx    *xtx.CXtx      //Gctx 全局腾讯操作接口对象
}

//RunCAuth 运行需要一直运行的协程
func (ca *CAuth) RunCAuth() {
	ca.Gcrs.UpdateSerRole()
	ca.Gctx.UpdateAccessToken()
	go ca.Gsm.GC()
	go ca.Gyzm.GC()
	go ca.Guk.RunUniKey()
}

//AddLogin 添加登陆信息
func (ca *CAuth) AddLogin(w http.ResponseWriter, tp, userid, phone, city string,
	rid []string) (sid string, err error) {
	switch tp {
	case "pc":
		sid, err = ca.Gsm.AddSession(userid, phone, city, rid)
		if err != nil {
			return
		}
		AddCookie(ca.pcCName, sid, w)
	case "wx":
		return
	case "h5":
		return
	}
	return
}

//DelLogin 注销登陆信息
func (ca *CAuth) DelLogin(r *http.Request, w http.ResponseWriter, tp string) {
	switch tp {
	case "pc":
		sid := GetCookie(ca.pcCName, r)
		ca.Gsm.DelSession(sid)
		DelCookie(ca.pcCName, sid, w)
	case "wx":
		return
	case "h5":
		return
	}
	return
}
