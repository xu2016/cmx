package xjhsql

import (
	_ "github.com/go-sql-driver/mysql"
	_ "gopkg.in/goracle.v2"
)

//CXSqlx 数据库访问类型
type CXSqlx struct {
	dbtype string
	db     string
	cmt    bool
}

//NewSQLX 创建一个数据库对象
func NewSQLX(sdbtype, sdb string, scmt bool) (xdb CXSqlx) {
	xdb = CXSqlx{dbtype: sdbtype, db: sdb, cmt: scmt}
	return
}
