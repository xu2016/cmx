package xsql

import (
	_ "github.com/go-sql-driver/mysql"
	_ "gopkg.in/goracle.v2"
)

//CXSql 数据库访问类型
type CXSql struct {
	dbtype string
	db     string
	cmt    bool
}

//NewSQLX 创建一个数据库对象
func NewSQL(sdbtype, sdb string, scmt bool) (xdb CXSql) {
	xdb = CXSql{dbtype: sdbtype, db: sdb, cmt: scmt}
	return
}
