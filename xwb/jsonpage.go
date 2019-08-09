package xwb

import (
	"encoding/json"
	"net/http"
)

//JSONPage 返回Json格式的页面
func JSONPage(cnt interface{}, w http.ResponseWriter) {
	data, _ := json.Marshal(cnt)
	w.Header().Set("x-frame-options", "SAMEORIGIN")
	w.Header().Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
