package xauth

import (
	"cmx/xcm"
	"cmx/xwb"
	"net/http"
)

//GCA 全局权限设置对象
var GCA *CAuth

//NewCAuth 创建一个全局权限设置对象
func NewCAuth(dbstr, dbtype string, cnt int, pcTime, wxTime, h5Time int64,
	pcCName, wxCName, h5CName string) *CAuth {
	return &CAuth{pcCName: pcCName, wxCName: wxCName, h5CName: h5CName, dbstr: dbstr, dbtype: dbtype,
		Gsm: NewSManager(pcTime), Gyzm: NewCyzm(3600), Guk: NewUniKey(3), Gcrs: NewCRoleSer(dbstr, dbtype)}
}

//CAuth 权限设置类
type CAuth struct {
	pcCName string
	wxCName string
	h5CName string
	dbstr   string
	dbtype  string
	Gsm     *SManager //Gpcsm 全局PC端Session管理器
	Gyzm    *Cyzm     //Gyzm 全局验证码管理器
	Guk     *CUniKey  //Guk 全局唯一key管理器
	Gcrs    *CRoleSer //Gcrs 全局角色服务对应关系管理对象
}

//RunCAuth 运行需要一直运行的协程
func (ca *CAuth) RunCAuth() {
	ca.Gcrs.UpdateSerRole()
	go ca.Gsm.GC()
	go ca.Gyzm.GC()
	go ca.Guk.runUniKey()
}

//IsLogin 判断是否登陆且有访问权限
func (ca *CAuth) IsLogin(r *http.Request, tp string) (sid string, b bool) {
	switch tp {
	case "pc":
		sid = GetCookie(ca.pcCName, r)
		roles := ca.Gsm.GetUserRoles(sid)
		serKey := xcm.GetMD5(xwb.GetURL(r) + r.FormValue("ctype"))
		b = ca.Gcrs.IsAllow(serKey, roles)
	case "wx":
		b = true
		return
	case "h5":
		b = true
		return
	}
	return
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
