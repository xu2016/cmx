package xwb

import "net/http"

//SetRouter 设置页面路由，furl为父url，形式如/test，Rt的string部分为/...
func SetRouter(furl string, Rt map[string]func(http.ResponseWriter, *http.Request)) {
	for k, v := range Rt {
		http.HandleFunc(furl+k, v)
	}
}
