package xauth

import "net/http"

//IsLogin 判断是否登陆且有访问权限
func (ca *CAuth) IsLogin(r *http.Request, tp string) (sid string, b bool) {
	switch tp {
	case "pc":
		sid = GetCookie(ca.pcCName, r)
		if sid == "" {
			return
		}
		uroles := ca.Gsm.GetUserRoles(sid)
		if len(uroles) == 0 {
			return
		}
		sroles, ok := ca.Gcrs.Query(r)
		if !ok {
			b = true
			return
		}
		b = isAllow(uroles, sroles)
		b = true
	case "wx":
		b = true
		return
	case "h5":
		b = true
		return
	}
	return
}

func isAllow(urs, srs []string) (allow bool) {
	for _, ur := range urs {
		for _, sr := range srs {
			if ur == sr {
				allow = true
			}
		}
	}
	return
}
