package xwb

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//EXPost post方法提交
func EXPost(cstr map[string]string, urlstr string) (body []byte, err error) {
	urlValues := url.Values{}
	for k, v := range cstr {
		urlValues.Add(k, v)
	}
	resp, err := http.PostForm(urlstr, urlValues)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

/*FXPost post方法提交,
tp："json"、"xml"，如果tp不是"json"或"xml"，rp将获取不到值，但仍然会返回解析后的[]byte类型body。
rp:如果返回的是"json"或"xml"，请填写需要解析成的结构体或map的指针类型(为nil时，rp无数据返回)。
body:返回resp.Body
*/
func FXPost(cstr map[string]string, urlstr string, tp string, rp interface{}) (body []byte, err error) {
	urlValues := url.Values{}
	for k, v := range cstr {
		urlValues.Add(k, v)
	}
	resp, err := http.PostForm(urlstr, urlValues)
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

/*PXPost post方法提交
contentType:为要提交的内容类型，比如，"application/json;charset=utf-8"
valStr：提交的内容，如果是"application/json;charset=utf-8"，就是一个json格式的字符串
tp："json"、"xml"，如果tp不是"json"或"xml"，rp将获取不到值，但仍然会返回解析后的[]byte类型body。
rp:如果返回的是"json"或"xml"，请填写需要解析成的结构体或map的指针类型(为nil时，rp无数据返回)。
body:返回resp.Body
*/
func PXPost(urlStr, contentType, valStr string, tp string, rp interface{}) (body []byte, err error) {
	resp, err := http.Post(urlStr, contentType, strings.NewReader(valStr))
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

/*NXPost post方法提交
contentType:为要提交的内容类型，比如，"application/json;charset=utf-8"
valStr：提交的内容，如果是"application/json;charset=utf-8"，就是一个json格式的字符串
tp："json"、"xml"，如果tp不是"json"或"xml"，rp将获取不到值，但仍然会返回解析后的[]byte类型body。
rp:如果返回的是"json"或"xml"，请填写需要解析成的结构体或map的指针类型(为nil时，rp无数据返回)。
body:返回resp.Body
*/
func NXPost(urlStr, contentType, valStr string, tp string, rp interface{}) (body []byte, err error) {
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(valStr))
	req.Header.Set("Content-Type", contentType)
	client := &http.Client{}
	resp, err := client.Do(req)
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
