package xtx

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

/*Gctx ...
腾讯相关访问接口
*/
var Gctx *CXtx

//NewCXtx ..
func NewCXtx(txid map[string]TxID, wpname string) *CXtx {
	return &CXtx{txid: txid, wpname: wpname, accessToken: "", expiresIn: 3600}
}

//CXtx 腾讯相关访问接口类
type CXtx struct {
	lock        sync.RWMutex
	txid        map[string]TxID
	wpname      string //公众号名称
	accessToken string
	expiresIn   int64
}

//TxID 腾讯公众号、小程序等的appid和appsecret结构体
type TxID struct {
	Appid  string
	Secret string
}

//UpdateAccessToken 定时更新腾讯公众号、小程序等的AccessToken
func (tx *CXtx) UpdateAccessToken() {
	t := time.NewTicker(time.Duration(tx.expiresIn-60) * time.Second)
	for {
		select {
		case <-t.C:
			tx.SetAccessToken()
		}
	}
}

type tokenJSON struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

//SetAccessToken 获取腾讯公众号、小程序等的AccessToken
func (tx *CXtx) SetAccessToken() {
	url := `https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=` + tx.txid[tx.wpname].Appid + `&secret=` + tx.txid[tx.wpname].Secret
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	var wx tokenJSON
	err = json.Unmarshal([]byte(body), &wx)
	if err != nil {
		log.Println(err)
		return
	}
	tx.lock.Lock()
	defer tx.lock.Unlock()
	log.Println(wx)
	if wx.Errcode == 0 {
		tx.accessToken = wx.AccessToken
		tx.expiresIn = wx.ExpiresIn
	} else {
		log.Println(wx.Errmsg)
	}
}

//GetAccessToken 获取腾讯公众号、小程序等的AccessToken
func (tx *CXtx) GetAccessToken() string {
	tx.lock.RLock()
	defer tx.lock.RUnlock()
	return tx.accessToken
}

type sendJSON struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

//SetMsgToUserXcx 小程序发送消息给用户
func (tx *CXtx) SetMsgToUserXcx(touser, templateid, page, formid, emphasiskeyword string, data map[string]string) (err error) {
	datas := ``
	n := len(data)
	for i := 0; i < n; i++ {
		k := `keyword` + strconv.Itoa(i+1) + ``
		v := data[k]
		if i != 0 {
			datas += `,`
		}
		datas += `"` + k + `":{"value": "` + v + `"}`
	}
	jsonStr := []byte(`{
							"touser": "` + touser + `",
							"template_id": "` + templateid + `",
							"page": "` + page + `",
							"form_id": "` + formid + `",
							"data": {` + datas + `},
							"emphasis_keyword": "` + emphasiskeyword + `"
						}`)
	tk := tx.GetAccessToken()
	url := `https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=` + tk
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	wx := &sendJSON{}
	json.Unmarshal([]byte(body), &wx)
	//log.Println(wx)
	if wx.Errcode != 0 {
		err = errors.New(wx.Errmsg)
		return
	}
	err = nil
	return
}
