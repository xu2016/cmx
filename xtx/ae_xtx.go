package xtx

/*Gctx ...
腾讯相关访问接口
*/
var Gctx *CXtx

//NewCXtx ..
func NewCXtx(txids map[string]TxID, tkname []string) *CXtx {
	tks := make(map[string]TxToken)
	for _, v := range tkname {
		tk := TxToken{AccessToken: "", ExpiresIn: 0}
		tks[v] = tk
	}
	return &CXtx{txids: txids, tokens: tks}
}

type txRJSON struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
