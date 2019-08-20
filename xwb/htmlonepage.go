package xwb

import (
	"html/template"
	"log"
	"net/http"
)

//HTMLOnePage 返回Html格式的页面,path+`/`+name
func HTMLOnePage(w http.ResponseWriter, path, name string) {
	t, _ := template.ParseFiles(path + `/` + name)
	w.Header().Set("x-frame-options", "SAMEORIGIN")
	w.Header().Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	w.WriteHeader(http.StatusOK)
	err := t.Execute(w, nil)
	if err != nil {
		log.Println("index:", err)
	}
}
