package xwb

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//EXPost post方法提交
func EXPost(cstr map[string]string, urlstr string) (body []byte, err error) {
	urlValues := url.Values{}
	for k, v := range cstr {
		urlValues.Add(k, v)
	}
	resp, err := http.PostForm(urlstr, urlValues)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return
}
