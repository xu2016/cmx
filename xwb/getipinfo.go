package xwb

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//IPInfo 返回JSON结构
type IPInfo struct {
	Code int          `json:"code"`
	Data TAOBAOIPJSON `json:"data"`
}

//TAOBAOIPJSON 淘宝接口返回的JSON
type TAOBAOIPJSON struct {
	Country   string `json:"country"`
	CountryID string `json:"country_id"`
	Area      string `json:"area"`
	AreaID    string `json:"area_id"`
	Region    string `json:"region"`
	RegionID  string `json:"region_id"`
	City      string `json:"city"`
	CityID    string `json:"city_id"`
	Isp       string `json:"isp"`
}

/*GetIPInfo 通过淘宝接口根据公网ip获取国家运营商等信息
**接口： http://ip.taobao.com/service/getIpInfo.php?ip=
 */
func GetIPInfo(ip string) *IPInfo {
	url := "http://ip.taobao.com/service/getIpInfo.php?ip="
	url += ip
	resp, err := http.Get(url)
	if err != nil {
		log.Println("xcm.GetIPInfo Get Error:", err)
		return nil
	}
	defer resp.Body.Close()
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("xcm.GetIPInfo ReadAll Error:", err)
		return nil
	}
	var result IPInfo
	if err := json.Unmarshal(out, &result); err != nil {
		log.Println("xcm.GetIPInfo Unmarshal Error:", err)
		return nil
	}
	return &result
}
