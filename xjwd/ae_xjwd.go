package xjwd

import "math"

//ER 地球半径
const ER = 6378137.0 // 6371000

//Rad ...
const Rad = math.Pi / 180.0

//ARad ...
const ARad = 180.0 / math.Pi

//EarthDistance 地球上两个经纬度之间的直线距离,单位：米
func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	lat1 = lat1 * Rad
	lat2 = lat2 * Rad
	theta := lng2*Rad - lng1*Rad
	return math.Floor(math.Acos(math.Sin(lat1)*math.Sin(lat2)+math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))*ER + 0.5)
}

//LngLatRange 获取距离经纬度为(lng,lat)的点距离为d的圆的外接正方形上下经纬度边界
func LngLatRange(lat, lng, d float64) (minLat, maxLat, maxLng, minLng float64) {
	rangex := ARad * d / 6372797.0 //里面的 d就代表搜索 dm 之内，单位m
	lngR := rangex / math.Cos(lat*Rad)
	maxLat = lat + rangex //最大纬度
	minLat = lat - rangex //最小纬度
	maxLng = lng + lngR   //最大经度
	minLng = lng - lngR   //最小经度
	return
}

const xpi = 3.14159265358979324 * 3000.0 / 180.0
const pi = 3.1415926535897932384626
const a = 6378245.0
const ee = 0.00669342162296594323

/*Bd09ToGcj02 百度坐标系 (BD-09) -> 火星坐标系 (GCJ-02)
 * 即 百度 转 谷歌、高德
 */
func Bd09ToGcj02(blng, blat float64) (glng, glat float64) {
	x := blng - 0.0065
	y := blat - 0.006
	z := math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*xpi)
	theta := math.Atan2(y, x) - 0.000003*math.Cos(x*xpi)
	glng = z * math.Cos(theta)
	glat = z * math.Sin(theta)
	return
}

/*Gcj02ToBd09 火星坐标系 (GCJ-02) ->百度坐标系 (BD-09)
 * 即谷歌、高德 转 百度
 */
func Gcj02ToBd09(glng, glat float64) (blng, blat float64) {
	z := math.Sqrt(glng*glng+glat*glat) + 0.00002*math.Sin(glat*xpi)
	theta := math.Atan2(glat, glng) + 0.000003*math.Cos(glng*xpi)
	blng = z*math.Cos(theta) + 0.0065
	blat = z*math.Sin(theta) + 0.006
	return
}
