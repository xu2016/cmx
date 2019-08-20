package xwb

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

//UnJSONForm 将浏览器或客户端返回过来的JSON数据解析到val中
func UnJSONForm(r *http.Request, val interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&val)
}

//UnXMLForm 将浏览器或客户端返回过来的xml数据解析到val中
func UnXMLForm(r *http.Request, val interface{}) error {
	decoder := xml.NewDecoder(r.Body)
	return decoder.Decode(&val)
}
