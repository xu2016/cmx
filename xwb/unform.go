package xwb

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
)

//UnForm 将浏览器或客户端返回过来的JSON或xml数据解析到val中
func UnForm(r *http.Request, val interface{}, tp string) error {
	if r.Body == nil {
		return errors.New(`request body is empty`)
	}
	var err error
	if tp == `json` {
		err = json.NewDecoder(r.Body).Decode(&val)
	} else if tp == `xml` {
		err = xml.NewDecoder(r.Body).Decode(&val)
	}
	if err != nil {
		return err
	}
	return nil
}
