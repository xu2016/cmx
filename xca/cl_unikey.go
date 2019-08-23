package xca

import (
	"cmx/xcm"
	"errors"
	"strings"
	"time"
)

var keyStr = [36]string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z"}

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

/*RunUniKey 新建一个UniKey生成器,每个生成器都需要使用go新建一个协程
cnt:表示生成的随机字符的位数，通过位数生成RN的最大的数
*/
func (cuk *CUniKey) RunUniKey() {
	sunix := time.Now().Unix()
	max := int64(0)
	x := int64(1)
	for i := 0; i < cuk.cnt; i++ {
		max += x * 35
		x *= 36
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
类型（位数具体确定）+年（2位）月（1位）日（1位）时（1位）分（1位）秒（1位）+编码位(1位)+随机位（位数具体由生成NewUniKey确定）
key的总长度:len(idtype)+8+len(max)
编码位(1位)：2*2*2*2*分位+2*2*2*分位+年位
	年位：年/1000(年位为0-7之间)
	分位：[0,29]为0,[30,59]为1。
	秒位：[0,29]为0,[30,59]为1。
max:
1位：[0,35]
2位：[0,35+35*36]=[0,1295]
3位：[0,35+35*36+35*36*36]=[0,46655]
4位：[0,35+35*36+35*36*36+35*36*36*36]=[0,1679615]
5位：[0,35+35*36+35*36*36+35*36*36*36+35*36*36*36*36]=[0,60466175]
6位：[0,35+35*36+35*36*36+35*36*36*36+35*36*36*36*36+35*36*36*36*36*36]=[0,2176782335]
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
		x := rnx.RN % 36
		rstr = keyStr[x] + rstr
		rnx.RN /= 36
	}
	tm := time.Unix(rnx.USecond, 0)
	yy, mm, dd, hh, ii, ss := xcm.GetTime(tm)
	yb := yy / 1000
	if yb > 7 {
		err = errors.New("year error")
		return
	}
	yy = yy % 1000
	ib := 0
	sb := 0
	if ii > 29 {
		ii -= 30
		ib = 1
	}
	if ss > 29 {
		ss -= 30
		sb = 1
	}
	bmb := 2*2*2*2*ib + 2*2*2*sb + yb
	id = strings.ToUpper(idHeader) + keyStr[yy/36] + keyStr[yy%36] + keyStr[mm] + keyStr[dd] +
		keyStr[hh] + keyStr[ii] + keyStr[ss] + keyStr[bmb] + rstr
	return
}
