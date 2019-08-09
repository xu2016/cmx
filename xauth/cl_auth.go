package xauth

import (
	"cmx/xcm"
	"cmx/xwb"
	"errors"
	"net/http"
)

//GCA 全局权限设置对象
var GCA *CAuth

//NewCAuth 创建一个全局权限设置对象
func NewCAuth(dbstr, dbtype string, uks, sessionTime map[string]int, cookiesName map[string]string) (ca *CAuth, err error) {
	ca = &CAuth{pcCookieName: cookiesName["pc"], wxCookieName: cookiesName["wx"], h5CookieName: cookiesName["h5"]}
	ca.Guk = NewUniKey(uks)
	if v, ok := sessionTime["pc"]; ok {
		ca.Gpcsm = NewSessionManager(int64(v))
	} else {
		err = errors.New(`sessionTime["pc"] not defined`)
	}
	ca.Gcrs = NewCRoleSer(dbstr, dbtype)
	return
}

//CAuth 权限设置类
type CAuth struct {
	pcCookieName string
	wxCookieName string
	h5CookieName string
	Gpcsm        *SManager //Gpcsm 全局PC端Session管理器
	Gyzm         *Cyzm     //Gyzm 全局验证码管理器
	Guk          *CUniKey  //Guk 全局唯一key管理器
	Gcrs         *CRoleSer //Gcrs 全局角色服务对应关系管理对象
}

//RunCAuth 运行需要一直运行的协程
func (ca *CAuth) RunCAuth() {
	ca.Gcrs.UpdateSerRole()
	go ca.Gpcsm.GC()
	go ca.Gyzm.GC()
}

//IsLogin 判断是否登陆且有访问权限
func (ca *CAuth) IsLogin(r *http.Request, tp string) (sid string, b bool) {
	switch tp {
	case "pc":
		sid = GetCookie(ca.pcCookieName, r)
		roles := ca.Gpcsm.GetUserRoles(sid)
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
func (ca *CAuth) AddLogin(w http.ResponseWriter, tp, cookieName, userid, phone, city string,
	rid []string) (sid string, err error) {
	switch tp {
	case "pc":
		sid, err = ca.Gpcsm.AddSession(userid, phone, city, rid)
		if err != nil {
			return
		}
		AddCookie(cookieName, sid, w)
	case "wx":
		return
	case "h5":
		return
	}
	return
}

//DelLogin 注销登陆信息
func (ca *CAuth) DelLogin(r *http.Request, w http.ResponseWriter, tp, cookieName string) {
	switch tp {
	case "pc":
		sid := GetCookie(cookieName, r)
		ca.Gpcsm.DelSession(sid)
		DelCookie(cookieName, sid, w)
	case "wx":
		return
	case "h5":
		return
	}
	return
}
