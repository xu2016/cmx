package xwb

/*
web操作相关功能封装
*/

//RJSON 仅仅返回成功和失败的JOSN
type RJSON struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//RJSONData 返回JOSN，包括数据
type RJSONData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
