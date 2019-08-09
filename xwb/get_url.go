package xwb

import "net/http"

//GetURL 获取访问的URL
func GetURL(r *http.Request) string {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	return scheme + r.Host + r.URL.Path
}
