package xjhcm

import (
	"os"
	"regexp"
	"strconv"
)

//IsPhone 判断num是不是手机号码
func IsPhone(num string) bool {
	rstr, err := regexp.Compile(`^(13|14|15|16|17|18|19)[0-9]{9}$`)
	if err != nil {
		return false
	}
	return rstr.MatchString(num)
}

//IsAlphaOrNum 判断str是内全部是0-9a-zA-Z，且长度为min到max之间
func IsAlphaOrNum(str string, min, max int) bool {
	rstr, err := regexp.Compile(`^[0-9a-zA-Z]{` + strconv.Itoa(min) + `,` + strconv.Itoa(max) + `}+$`)
	if err != nil {
		return false
	}
	return rstr.MatchString(str)
}

//PathExists 判断路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
