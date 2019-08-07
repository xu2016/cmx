package xjhauth

//GCSer 全局服务管理对象
var GCSer *CSer

//NewCSer 创建一个服务管理对象
func NewCSer(dbstr string, dbtype string, dbcommit string) *CSer {
	return &CSer{dbstr: dbstr, dbtype: dbtype, dbcommit: dbcommit}
}

//CSer 服务管理类
type CSer struct {
	dbstr    string //数据库连接信息
	dbtype   string //数据库类型"mysql"或"goracle"
	dbcommit string //insert和update是否要commit，"goracle"时为true
}

func (cs *CSer) Add() {

}
