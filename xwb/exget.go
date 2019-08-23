package xwb

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

/*EXGet Get方法提交,
tp："json"、"xml"，如果tp不是"json"或"xml"，rp将获取不到值，但仍然会返回解析后的[]byte类型body。
rp:如果返回的是"json"或"xml"，请填写需要解析成的结构体或map的指针类型(为nil时，rp无数据返回)。
body:返回resp.Body
*/
func EXGet(urlstr string, tp string, rp interface{}) (body []byte, err error) {
	resp, err := http.Get(urlstr)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if rp == nil && (tp == "json" || tp == "xml") {
		err = errors.New(`tp is ` + tp + ` but rp is nil`)
		return
	}
	switch tp {
	case `json`:
		err = json.Unmarshal(body, rp)
	case `xml`:
		err = xml.Unmarshal(body, rp)
	}
	return
}
