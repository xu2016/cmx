package xjwd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type addresstoBd09JSON struct {
	Status int             `json:"status"`
	Result dz2zbResultJSON `json:"result"`
}
type dz2zbResultJSON struct {
	Location locationJSON `json:"location"`
}
type locationJSON struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

//Dz2Bd09 把地址转换成BD09ll坐标系
//url=http://api.map.baidu.com/geocoder/v2/?address=北京市海淀区上地十街10号&output=json&ak=您的ak&callback=showLocation
func Dz2Bd09(address, key string) (bd09lng, bd09lat float64, err error) {
	var bd addresstoBd09JSON
	bd.Status = 1
	url := fmt.Sprintf("http://api.map.baidu.com/geocoding/v3/?address=%s&output=json&ak=%s&callback=showLocation", address, key)
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
	if bd.Status == 0 {
		bd09lng = bd.Result.Location.Lng
		bd09lat = bd.Result.Location.Lat
		err = nil
	}
	return
}
