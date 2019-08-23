package xtx

import (
	"cmx/xwb"
	"errors"
	"log"
	"strconv"
)

//SetMsgToUserXcx 小程序发送消息给用户
func (tx *CXtx) SetMsgToUserXcx(touser, templateid, page, formid, emphasiskeyword, tokenName string, data map[string]string) (err error) {
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
	jsonStr := `{
					"touser": "` + touser + `",
					"template_id": "` + templateid + `",
					"page": "` + page + `",
					"form_id": "` + formid + `",
					"data": {` + datas + `},
					"emphasis_keyword": "` + emphasiskeyword + `"
				}`
	tk := tx.GetAccessToken(tokenName)
	url := `https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=` + tk
	wx := txRJSON{}
	_, err = xwb.PXPost(url, "application/json;charset=utf-8", jsonStr, "json", &wx)
	if err != nil {
		log.Println(err)
		return
	}
	//log.Println(wx)
	if wx.Errcode != 0 {
		err = errors.New(wx.Errmsg)
		return
	}
	err = nil
	return
}
