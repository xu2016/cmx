package xjwd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
{
	"status":0,
	"result":
	{
		"location":
		{
			"lng":113.43183587093087,"lat":22.513980292596366
		},
		"formatted_address":"广东省中山市中山市市辖区博爱六路",
		"business":"",
		"addressComponent":
		{
			"country":"中国",
			"country_code":0,
			"country_code_iso":
			"CHN",
			"country_code_iso2":"CN",
			"province":"广东省",
			"city":"中山市",
			"city_level":2,
			"district":"中山市市辖区",
			"town":"",
			"adcode":"442000",
			"street":"博爱六路",
			"street_number":"",
			"direction":"",
			"distance":""
		},
		"pois":[],
		"roads":[],
		"poiRegions":[],
		"sematic_description":"联通广场附近0米",
		"cityCode":187
	}
}
*/

type bd09toaddressJSON struct {
	Status int             `json:"status"`
	Result zb2dzResultJSON `json:"result"`
}
type zb2dzResultJSON struct {
	FormattedAddress   string `json:"formatted_address"`
	SematicDescription string `json:"sematic_description"`
}

//Bd09toDz 把BD09ll坐标系转换成地址
//url=http://api.map.baidu.com/geocoder/v2/?location=22.513980293547004,113.4318358709309&output=json&pois=0&ak=FgDPj4Ey2493stHqR6Ns2SiLCwD8VPqT
func Bd09toDz(lng, lat float64, key string) (FormattedAddress, sematicDescription string, err error) {
	var bd bd09toaddressJSON
	bd.Status = 1
	url := fmt.Sprintf(`http://api.map.baidu.com/reverse_geocoding/v3/?ak=%s&output=json&location=%.8f,%.8f`, key, lat, lng)
	log.Println(url)
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
		FormattedAddress = bd.Result.FormattedAddress
		sematicDescription = bd.Result.SematicDescription
		err = nil
	}
	return
}
