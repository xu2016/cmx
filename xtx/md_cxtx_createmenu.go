package xtx

type createMenuJSON struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

/*CreateMenu 腾讯公众号创建菜单


 */
func (tx *CXtx) CreateMenu(tokenName string) {
	// tk := tx.GetAccessToken(tokenName)
	// rp := &createMenuJSON{}
	// url := `https://api.weixin.qq.com/cgi-bin/menu/create?access_token=` + tk
	// jsonStr := `{}`

}

func createMenuJSONSTR() (jsonStr string, err error) {

	return
}
