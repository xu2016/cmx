package xcm

import (
	"io/ioutil"
	"os"
	"strings"
)

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

/*GetAllFile 获取pathname下所有后缀为suffix的文件，
  level是路径深度，pathname下为第1层，当level==-1时，遍历pathname下所有目录
  s为存放遍历得到的所有文件路径（包括文件名）
*/
func GetAllFile(pathname, suffix string, level int, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			if level > 1 || level == -1 {
				l := level
				if level > 1 {
					l--
				}
				fullDir := pathname + "/" + fi.Name()
				s, err = GetAllFile(fullDir, suffix, l, s)
				if err != nil {
					return s, err
				}
			}
		} else {
			fiName := fi.Name()
			if strings.HasSuffix(fiName, suffix) {
				fullName := pathname + "/" + fiName
				s = append(s, fullName)
			}
		}
	}
	return s, nil
}
