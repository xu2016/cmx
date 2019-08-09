package xauth

import "net/http"

var dbcommit = map[string]bool{"mysql": false, "goracle": true}
var CookiesName string

//AddCookie 添加cookie
func AddCookie(name, val string, w http.ResponseWriter) {
	cookie := http.Cookie{Name: name, Value: val, Path: "/", HttpOnly: true}
	http.SetCookie(w, &cookie)
}

//DelCookie 删除cookie
func DelCookie(name, val string, w http.ResponseWriter) {
	cookie := http.Cookie{Name: name, Value: val, Path: "/", HttpOnly: true, MaxAge: -1}
	http.SetCookie(w, &cookie)
}

//GetCookie 删除cookie
func GetCookie(name string, r *http.Request) string {
	cookie, err := r.Cookie(name)
	if err != nil {
		return ""
	}
	return cookie.Value
}
