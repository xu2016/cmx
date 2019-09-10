package xwb

import (
	"compress/gzip"
	"io/ioutil"
	"net/http"
	"strings"
)

/*GetStaticFile 获取浏览器需要的静态资源文件
header('Content-Type: text/html; charset=utf-8'); //网页编码
header('Content-Type: text/plain'); //纯文本格式
header('Content-Type: image/jpeg'); //JPG、JPEG
header('Content-Type: application/zip'); // ZIP文件
header('Content-Type: application/pdf'); // PDF文件
header('Content-Type: audio/mpeg'); // 音频文件
header('Content-type: video/mp4'); //mp4视频文件
header('Content-type: text/css'); //css文件
header('Content-type: text/javascript'); //js文件
header('Content-type: application/json'); //json
header('Content-type: application/pdf'); //pdf
header('Content-type: text/xml'); //xml
https://www.runoob.com/http/http-content-type.html
*/
var filetype = map[string]string{
	`css`:  `text/css`,
	`js`:   `text/javascript`,
	`txt`:  `text/plain`,
	`png`:  `image/png`,
	`jpg`:  `image/jpeg`,
	`jpeg`: `image/jpeg`,
	`ico`:  `image/x-icon`,
	`bmp`:  `application/x-bmp`,
	`mp4`:  `video/mp4`,
	`html`: `text/html`,
	`htm`:  `text/html`,
	`wav`:  `audio/wav`,
	`wma`:  `audio/x-ms-wma`,
	`xls`:  `application/vnd.ms-excel`,
	`mp3`:  `audio/mp3`,
	`midi`: `audio/mid`,
	`xlsx`: `application/vnd.openxmlformats-officedocument.spreadsheetml.sheet`,
}

//GetStaticFile 获取服务器上的静态文件
func GetStaticFile(w http.ResponseWriter, r *http.Request, rootpath string) {
	path := r.URL.Path
	pathStr := strings.Split(path, ".")
	v, ok := filetype[pathStr[len(pathStr)-1]]
	if !ok {
		w.Write([]byte(`<h1>页面正在更新。。。。。</h1>`))
		return
	}
	file, err := ioutil.ReadFile(rootpath + path)
	if err != nil {
		w.Write([]byte(`<h1>页面正在更新。。。。。</h1>`))
		return
	}
	w.Header().Set(`Content-type`, v)
	//开启gzip压缩
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Accept-Encoding", "gzip")
	gzw := gzip.NewWriter(w)
	//w.Write(file)
	defer gzw.Close()
	gzw.Write(file)
}
