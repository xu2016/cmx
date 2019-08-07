package xjhauth

func init() {
	GUniKeyNames = make(map[string]int, 0)
	GUniKeyNames[`cs`] = 1  //服务全局唯一key管理器
	GUniKeyNames[`cr`] = 1  //角色全局唯一key管理器
	GUniKeyNames[`cur`] = 1 //用户角色对应关系全局唯一key管理器
	GUniKeyNames[`csr`] = 1 //服务角色对应关系全局唯一key管理器
	GYzm = NewCyzm(600)
	go GYzm.GC()
}

//GUniKeyNames 不同全局唯一key管理器的名称和随机数位数
var GUniKeyNames map[string]int

//InitUniKey 初始化全局唯一key管理器
func InitUniKey(str map[string]int) {
	for k, v := range str {
		if
	}
}
