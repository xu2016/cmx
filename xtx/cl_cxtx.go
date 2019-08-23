package xtx

import (
	"cmx/xwb"
	"log"
	"sync"
	"time"
)

//CXtx 腾讯相关访问接口类
type CXtx struct {
	lock   sync.RWMutex
	txids  map[string]TxID
	tokens map[string]TxToken
}

//TxID 腾讯公众号、小程序等的appid和appsecret结构体
type TxID struct {
	Appid  string
	Secret string
}

//TxToken 腾讯AccessToken结构体
type TxToken struct {
	AccessToken string
	ExpiresIn   int64
}

//GetAccessToken 获取腾讯公众号、小程序等的AccessToken
func (tx *CXtx) GetAccessToken(tokenName string) string {
	tx.lock.RLock()
	defer tx.lock.RUnlock()
	tk, ok := tx.tokens[tokenName]
	if !ok {
		log.Println("Token " + tokenName + " not defined")
		return ""
	}
	return tk.AccessToken
}

//UpdateAccessToken 定时更新腾讯公众号、小程序等的AccessToken
func (tx *CXtx) UpdateAccessToken() {
	for k := range tx.tokens {
		tx.setAccessToken(k)
	}
	for k, v := range tx.tokens {
		go tx.updateAccessToken(k, v.ExpiresIn)
	}
}

type tokenJSON struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

//updateAccessToken 定时更新腾讯公众号、小程序等的AccessToken
func (tx *CXtx) updateAccessToken(tokenName string, times int64) {
	times = times - 60
	if times <= 0 {
		times = 6
	}
	t := time.NewTicker(time.Duration(times) * time.Second)
	for {
		select {
		case <-t.C:
			times = tx.setAccessToken(tokenName)
			t = time.NewTicker(time.Duration(times) * time.Second)
		}
	}
}

//setAccessToken 获取腾讯公众号、小程序等的AccessToken
func (tx *CXtx) setAccessToken(tokenName string) (times int64) {
	tx.lock.Lock()
	defer tx.lock.Unlock()
	txid, ok := tx.txids[tokenName]
	if !ok {
		log.Println("Token " + tokenName + " not defined")
		return
	}
	appid := txid.Appid
	secret := txid.Secret
	url := `https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=` + appid + `&secret=` + secret
	var wx tokenJSON
	_, err := xwb.EXGet(url, "json", &wx)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(wx)
	if wx.Errcode == 0 {
		tk := TxToken{AccessToken: wx.AccessToken, ExpiresIn: wx.ExpiresIn}
		tx.tokens[tokenName] = tk
		times = wx.ExpiresIn - 60
		if times <= 0 {
			times = 6
		}
	} else {
		log.Println(wx.Errmsg)
		times = 6
	}
	return
}
