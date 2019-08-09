package xsql

import (
	"database/sql"
	"errors"
)

/*QueryLine 查询一条记录
  parameter：
	qstr：查询语句，里面包含一些需要替换的值在val中,替换标识号取决于不同的数据库
		  MySQL：?,?,....,?
		  Oracle：:1,:2,.....,:n
		注意：在qstr中要处理好可能出现NULL的情况，比如在Oracle中，nvl(字段，为NULL是的值)
	val：要替换的值的slice
	col：需要取的字段的slice。
	coltype：需要取的字段所对应的数据类型，包括int8,uint8,int16,uint16,int32,uint32,int64,uint64,int,uint,float32,float64,bool,string
  return：返回结果是一个map[string]interface{}，请获取到数据后自行进行数据转换,操作错误err。
*/
func (xdb *CXSql) QueryLine(qstr string, val []interface{}, col []string, coltype []string) (gdata map[string]interface{}, err error) {
	db, err := sql.Open(xdb.dbtype, xdb.db)
	if err != nil {
		return
	}
	defer db.Close()
	if len(col) == 0 || len(col) != len(coltype) {
		err = errors.New("参数有误")
		return
	}
	sci := make([]interface{}, len(col))
	for i := range sci {
		switch coltype[i] {
		case `int8`:
			sci[i] = new(int8)
		case ` uint8`:
			sci[i] = new(uint8)
		case `int16`:
			sci[i] = new(int16)
		case `uint16`:
			sci[i] = new(uint16)
		case `int32`:
			sci[i] = new(int32)
		case `uint32`:
			sci[i] = new(uint32)
		case `int64`:
			sci[i] = new(int64)
		case `uint64`:
			sci[i] = new(uint64)
		case `int`:
			sci[i] = new(int)
		case `uint`:
			sci[i] = new(uint)
		case `float32`:
			sci[i] = new(float32)
		case `float64`:
			sci[i] = new(float64)
		case `string`:
			sci[i] = new(string)
		case `bool`:
			sci[i] = new(bool)
		default:
			sci[i] = new(string)
		}
	}
	row := db.QueryRow(qstr, val...)
	if err != nil {
		//log.Println("Query Error:", err)
		//log.Println("Query Sql:", qstr)
		return
	}
	err = row.Scan(sci...)
	if err != nil {
		//log.Println("Query Scan:", err)
		return
	}
	ds := make(map[string]interface{})
	for i, v := range sci {
		switch coltype[i] {
		case `int8`:
			ds[col[i]] = *v.(*int8)
		case ` uint8`:
			ds[col[i]] = *v.(*uint8)
		case `int16`:
			ds[col[i]] = *v.(*int16)
		case `uint16`:
			ds[col[i]] = *v.(*uint16)
		case `int32`:
			ds[col[i]] = *v.(*int32)
		case `uint32`:
			ds[col[i]] = *v.(*uint32)
		case `int64`:
			ds[col[i]] = *v.(*int64)
		case `uint64`:
			ds[col[i]] = *v.(*uint64)
		case `int`:
			ds[col[i]] = *v.(*int)
		case `uint`:
			ds[col[i]] = *v.(*uint)
		case `float32`:
			ds[col[i]] = *v.(*float32)
		case `float64`:
			ds[col[i]] = *v.(*float64)
		case `string`:
			ds[col[i]] = *v.(*string)
		case `bool`:
			ds[col[i]] = *v.(*bool)
		default:
			ds[col[i]] = *v.(*string)
		}
	}
	gdata = ds
	return
}

/*Query 查询多条记录
  parameter：
	qstr：查询语句，里面包含一些需要替换的值在val中,替换标识号取决于不同的数据库
		  MySQL：?,?,....,?
		  Oracle：:1,:2,.....,:n
		注意：在qstr中要处理好可能出现NULL的情况，比如在Oracle中，nvl(字段，为NULL是的值)
	val：要替换的值的slice
	col：需要取的字段的slice。
	coltype：需要取的字段所对应的数据类型，包括int8,uint8,int16,uint16,int32,uint32,int64,uint64,int,uint,float32,float64,bool,string
  return：返回结果是一个[]map[string]interface{}，请获取到数据后自行进行数据转换。
*/
func (xdb *CXSql) Query(qstr string, val []interface{}, col []string, coltype []string) (gdata []map[string]interface{}, err error) {
	db, err := sql.Open(xdb.dbtype, xdb.db)
	if err != nil {
		return
	}
	defer db.Close()
	if len(col) == 0 || len(col) != len(coltype) {
		err = errors.New("参数有误")
		return
	}
	sci := make([]interface{}, len(col))
	for i := range sci {
		switch coltype[i] {
		case `int8`:
			sci[i] = new(int8)
		case ` uint8`:
			sci[i] = new(uint8)
		case `int16`:
			sci[i] = new(int16)
		case `uint16`:
			sci[i] = new(uint16)
		case `int32`:
			sci[i] = new(int32)
		case `uint32`:
			sci[i] = new(uint32)
		case `int64`:
			sci[i] = new(int64)
		case `uint64`:
			sci[i] = new(uint64)
		case `int`:
			sci[i] = new(int)
		case `uint`:
			sci[i] = new(uint)
		case `float32`:
			sci[i] = new(float32)
		case `float64`:
			sci[i] = new(float64)
		case `string`:
			sci[i] = new(string)
		case `bool`:
			sci[i] = new(bool)
		default:
			sci[i] = new(string)
		}
	}
	rows, err := db.Query(qstr, val...)
	if err != nil {
		//log.Println("Query Error:", err)
		//log.Println("Query Sql:", qstr)
		return
	}
	for rows.Next() {
		err = rows.Scan(sci...)
		if err != nil {
			continue
		}
		ds := make(map[string]interface{})
		for i, v := range sci {
			switch coltype[i] {
			case `int8`:
				ds[col[i]] = *v.(*int8)
			case ` uint8`:
				ds[col[i]] = *v.(*uint8)
			case `int16`:
				ds[col[i]] = *v.(*int16)
			case `uint16`:
				ds[col[i]] = *v.(*uint16)
			case `int32`:
				ds[col[i]] = *v.(*int32)
			case `uint32`:
				ds[col[i]] = *v.(*uint32)
			case `int64`:
				ds[col[i]] = *v.(*int64)
			case `uint64`:
				ds[col[i]] = *v.(*uint64)
			case `int`:
				ds[col[i]] = *v.(*int)
			case `uint`:
				ds[col[i]] = *v.(*uint)
			case `float32`:
				ds[col[i]] = *v.(*float32)
			case `float64`:
				ds[col[i]] = *v.(*float64)
			case `string`:
				ds[col[i]] = *v.(*string)
			case `bool`:
				ds[col[i]] = *v.(*bool)
			default:
				ds[col[i]] = *v.(*string)
			}
		}
		gdata = append(gdata, ds)
	}
	return
}
