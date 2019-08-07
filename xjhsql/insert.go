package xjhsql

import (
	"database/sql"
)

/*Insertline 插入一条记录
  parameter：
	istr：查询语句，里面包含一些需要替换的值在val中,替换标识号取决于不同的数据库
		  MySQL：?,?,....,?
		  Oracle：:1,:2,.....,:n
	val：要替换的值的slice
  return：返回操作错误err。
*/
func (xdb *CXSqlx) Insertline(istr string, val []interface{}) (err error) {
	db, err := sql.Open(xdb.dbtype, xdb.db)
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec(istr, val...)
	if err != nil {
		//log.Println("Insertline Error:", err)
		//log.Println("Insertline Sql:", istr)
		return
	}
	if xdb.cmt {
		_, err = db.Exec(`commit`)
	}
	return
}
