package xtx

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type wxcodeInfo struct {
	Openid     string `json:"openid"`
	Sessionkey string `json:"session_key"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

/*GetOpenID 用于获取用户的微信openid
功能：用于获取用户的微信openid
参数：
	url：https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
		参数		必填	说明
		appid		是	小程序唯一标识
		secret		是	小程序的 app secret
		js_code		是	登录时获取的 code
		grant_type	是	填写为 authorization_code
返回：
	openid：成功返回用户微信openid，失败返回空字符串
	err:成功返回nil，失败返回错误
*/
func GetOpenID(wxAppID, wxAppSecret, wxcode string) (openid string, err error) {
	var wx wxcodeInfo
	url := `https://api.weixin.qq.com/sns/jscode2session?appid=` + wxAppID + `&secret=` + wxAppSecret + `&js_code=` + wxcode + `&grant_type=authorization_code`
	resp, err := http.Get(url)
	openid = "0"
	if err != nil {
		err = errors.New("xwb.GetOpenID Get提交失败:" + err.Error())
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.New("xwb.GetOpenID 读取HTML内容失败:" + err.Error())
		return
	}
	err = json.Unmarshal([]byte(body), &wx)
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
