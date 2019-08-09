package xcm

import (
	"errors"
	"math/rand"
	"time"
)

/*GetRandomInt 生成随机数字
[min,max)
*/
func GetRandomInt(min, max int) (rd int, err error) {
	err = nil
	if max < min {
		err = errors.New("max 必须大于 min")
		return
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rd = r.Intn(max-min-1) + min
	return
}

/*GetRandomSeedInt 生成随机数字
[min,max)
*/
func GetRandomSeedInt(min, max int, seed int64) (rd int, err error) {
	err = nil
	if max < min {
		err = errors.New("max 必须大于 min")
		return
	}
	r := rand.New(rand.NewSource(seed))
	rd = r.Intn(max-min-1) + min
	return
}

/*GetRandomString 生成随机字符串
slen:生成的随机数长度
stp:加密所选择的类型
    NSTR   = iota //数字字符串
	SDSTR         //小写字母字符串
	SUSTR         //大写字母字符串
	SASTR         //大写和小写字母字符串
	NSDSTR        //数字和小写字母字符串
	NSUSTR        //数字和大写字母字符串
	NSASTR        //数字大写和小写字母字符串
	KEYSTR        //数字大写和小写字母字符串(有序)
*/
func GetRandomString(slen int64, stp int) string {
	var mstr string
	switch stp {
	case NSTR:
		mstr = nstr
	case SDSTR:
		mstr = sdstr
	case SUSTR:
		mstr = sustr
	case SASTR:
		mstr = sastr
	case NSDSTR:
		mstr = nsdstr
	case NSUSTR:
		mstr = nsustr
	case KEYSTR:
		mstr = keystr
	default:
		mstr = nsastr
	}
	bytes := []byte(mstr)
	blen := len(bytes)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := int64(0); i < slen; i++ {
		result = append(result, bytes[r.Intn(blen)])
	}
	return string(result)
}
