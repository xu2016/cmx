package xjhsql

import (
	"database/sql"
)

/*Deleteline 删除一条记录
  parameter：
	dstr：查询语句，里面包含一些需要替换的值在val中,替换标识号取决于不同的数据库
		  MySQL：?,?,....,?
		  Oracle：:1,:2,.....,:n
	val：要替换的值的slice
  return：返回操作错误err。
*/
func (xdb *CXSqlx) Deleteline(dstr string, val []interface{}) (err error) {
	db, err := sql.Open(xdb.dbtype, xdb.db)
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec(dstr, val...)
	if err != nil {
		//log.Println("Deleteline Error:", err)
		//log.Println("Deleteline Sql:", dstr)
		return
	}
	if xdb.cmt {
		_, err = db.Exec(`commit`)
	}
	return
}
