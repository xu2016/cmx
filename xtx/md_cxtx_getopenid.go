package xtx

import (
	"cmx/xwb"
	"errors"
	"log"
)

type wxOpenIDcodeInfo struct {
	Openid     string `json:"openid"`
	Sessionkey string `json:"session_key"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

//GetOpenID 获取腾讯公众号、小程序等的openid
func (tx *CXtx) GetOpenID(wxname, wxcode string) (openid string, err error) {
	var wx wxOpenIDcodeInfo
	tx.lock.RLock()
	appid := tx.txids[wxname].Appid
	secret := tx.txids[wxname].Secret
	tx.lock.RUnlock()
	url := `https://api.weixin.qq.com/sns/jscode2session?appid=` + appid + `&secret=` + secret + `&js_code=` + wxcode + `&grant_type=authorization_code`
	_, err = xwb.EXGet(url, "json", &wx)
	if err != nil {
		log.Println(err)
		return
	}
	if wx.Openid == "" {
		err = errors.New("xwb.GetOpenID wx.Errmsg:" + wx.Errmsg)
		log.Println(err)
		openid = "0"
		return
	}
	openid = wx.Openid
	err = nil
	return
}
