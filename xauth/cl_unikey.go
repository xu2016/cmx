package xauth

import (
	"cmx/xcm"
	"errors"
	"time"
)

var keyStr = [62]string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d",
	"e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
	"o", "p", "q", "r", "s", "t", "u", "v", "w", "x",
	"y", "z"}

//CUniKey 唯一key管理器
type CUniKey struct {
	uks chan uniKey
	cnt int
}

//UniKey 唯一key定义
type uniKey struct {
	RN      int64
	USecond int64
	Cnt     int
}

/*NewUniKey 创建一个Session管理器
uknames:每个key键表示一个UniKey生成器，value表示生成器的生成的随机数的位数。
*/
func NewUniKey(cnt int) *CUniKey {
	return &CUniKey{uks: make(chan uniKey), cnt: cnt}
}

/*runUniKey 新建一个UniKey生成器,每个生成器都需要使用go新建一个协程
cnt:表示生成的随机字符的位数，通过位数生成RN的最大的数
*/
func (cuk *CUniKey) runUniKey() {
	sunix := time.Now().Unix()
	max := int64(0)
	x := int64(1)
	for i := 0; i < cuk.cnt; i++ {
		max += x * 61
		x *= 62
	}
	for i := int64(0); ; i++ {
		if sunix != time.Now().Unix() {
			sunix = time.Now().Unix()
			i = 0
		}
		if i <= max {
			cuk.uks <- uniKey{RN: i, USecond: sunix}
		}
	}
}

/*GetUniKey 获取生成的唯一的编号，编号格式如下：
类型（位数具体确定）+年（2位）月（1位）日（1位）时（1位）分（1位）秒（1位）+随机位（位数具体由生成NewUniKey确定）
key的总长度:len(idtype)+7+len(max)
1位：[0,61]
2位：[0,61+61*62]=[0,3843]
3位：[0,61+61*62+61*62*62]=[0,238327]
4位：[0,61+61*62+61*62*62+61*62*62*62]=[0,14776335]
5位：[0,61+61*62+61*62*62+61*62*62*62+61*62*62*62*62]=[0,916132831]
...
*/
func (cuk *CUniKey) GetUniKey(idHeader string) (id string, err error) {
	rnx, ok := <-cuk.uks
	if !ok {
		err = errors.New("get unikey error")
		return
	}
	rstr := ""
	for i := 0; i < cuk.cnt; i++ {
		x := rnx.RN % 62
		rstr = keyStr[x] + rstr
		rnx.RN /= 62
	}
	tm := time.Unix(rnx.USecond, 0)
	yy, mm, dd, hh, ii, ss := xcm.GetTime(tm)
	id = idHeader + keyStr[yy/62] + keyStr[yy%62] + keyStr[mm] + keyStr[dd] +
		keyStr[hh] + keyStr[ii] + keyStr[ss] + rstr
	return
}
