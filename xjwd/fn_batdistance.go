package xjwd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type batDistanceJSON struct {
	Status  int      `json:"status"`
	Result  []result `json:"result"`
	Message string   `json:"message"`
}
type result struct {
	Distance rType `json:"distance"`
	Duration rType `json:"duration"`
}
type rType struct {
	Text  string  `json:"text"`
	Value float64 `json:"value"`
}

//BatDistance 计算两点的步行距离
//url=http://api.map.baidu.com/routematrix/v2/walking?output=json&origins=40.45,116.34&destinations=40.34,116.45&ak=FgDPj4Ey2493stHqR6Ns2SiLCwD8VPqT
func BatDistance(src, des []float64, key string) (distance, duration []float64, err error) {
	var bd batDistanceJSON
	bd.Status = 1
	url = url + "?output=json&origins=" + strconv.FormatFloat(slat, 'f', -1, 64) + "," + strconv.FormatFloat(slng, 'f', -1, 64) + "&destinations=" + strconv.FormatFloat(dlat, 'f', -1, 64) + "," + strconv.FormatFloat(dlng, 'f', -1, 64) + "&ak=" + key
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(body), &bd)
	//log.Println(err, bd)
	if bd.Status == 0 {
		distance = bd.Result[0]["distance"].Value
		err = nil
	}
	return
}
