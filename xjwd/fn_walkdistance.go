package xjwd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type walkDistanceJSON struct {
	Status  int                 `json:"status"`
	Result  []map[string]inJSON `json:"result"`
	Message string              `json:"message"`
}
type inJSON struct {
	Text  string  `json:"text"`
	Value float64 `json:"value"`
}

//WalkDistance 计算两点的步行距离
//url=http://api.map.baidu.com/routematrix/v2/walking?output=json&origins=40.45,116.34&destinations=40.34,116.45&ak=FgDPj4Ey2493stHqR6Ns2SiLCwD8VPqT
func WalkDistance(url string, slng, slat, dlng, dlat float64, key string) (distance float64, err error) {
	var bd walkDistanceJSON
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
