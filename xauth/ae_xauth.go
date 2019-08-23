package xauth

import (
	"cmx/xca"
	"cmx/xtx"
)

var dbcommit = map[string]bool{"mysql": false, "goracle": true}

//GCA 全局权限设置对象
var GCA *CAuth

//NewCAuth 创建一个全局权限设置对象
func NewCAuth(dbstr, dbtype string, cnt int, pcTime, wxTime, h5Time int64, pcCName, wxCName, h5CName string,
	txids map[string]xtx.TxID, tkname []string) *CAuth {
	return &CAuth{pcCName: pcCName, wxCName: wxCName, h5CName: h5CName, dbstr: dbstr, dbtype: dbtype,
		Gsm:  xca.NewSManager(pcTime),
		Gyzm: xca.NewCyzmCache(3600),
		Guk:  xca.NewUniKey(3),
		Gcrs: xca.NewCRoleSer(dbstr, dbtype, dbcommit[dbtype]),
		Gctx: xtx.NewCXtx(txids, tkname)}
}
