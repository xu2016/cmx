package xtx

import (
	"bytes"
	"cmx/xauth"
	"cmx/xcm"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

//YzmInfo 返回验证码信息
type YzmInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Yzm  string `json:"yzm"`
}

//YzmJkInfo 返回验证码信息
type YzmJkInfo struct {
	Result int    `json:"result"`
	Errmsg string `json:"errmsg"`
	Ext    string `json:"ext"`
	Fee    int    `json:"fee"`
	Sid    string `json:"sid"`
}

/*SendTxYzm 发送腾讯平台短信
**phone:要发送的号码
**sdkappid:腾讯短信平台SDK APPID
**appkey:腾讯短信平台APP KEY
**params:腾讯短信平台短信模板中需要修改的内容，具体格式如下：
**     "内容1",
**     "内容2",
**     "内容3",
**     "内容4",
**     "....",
**     "内容n"
**sign:短信标题，腾讯短信平台短信模板的标题内容。
**tplid:腾讯短信平台短信模板的模板ID。
**random:随机验证码，可以为空字符串。
 */
func SendTxYzm(phone, sdkappid, appkey, params, sign, tplid, random string) (err error) {
	tm := strconv.FormatInt(time.Now().Unix(), 10)
	str := `appkey=` + appkey + `&random=` + random + `&time=` + tm + `&mobile=` + phone
	sig := xcm.GetSHA256(str)
	jsonStr := []byte(`{
		"ext": "",
		"extend": "",
		"params": [` + params + `],
		"sig": "` + sig + `",
		"sign": "` + sign + `",
		"tel": {
			"mobile": "` + phone + `",
			"nationcode": "86"
		},
		"time": ` + tm + `,
		"tpl_id": ` + tplid + `
	}`)
	xauth.GYzm.SetYzm(random, phone)
	url := `https://yun.tim.qq.com/v5/tlssmssvr/sendsms?sdkappid=` + sdkappid + `&random=` + random
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(phone+":", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(phone+":", err)
		return
	}
	wx := &YzmJkInfo{}
	json.Unmarshal([]byte(body), &wx)
	if wx.Result != 0 {
		log.Println(phone+":", wx.Errmsg)
		err = errors.New(wx.Errmsg)
		return
	}
	return
}
