package xjhcm

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
