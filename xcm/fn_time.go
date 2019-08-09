package xcm

import "time"

//GetNowTime 返回当前时间的Year, Month, Day, Hour, Minute, Second
func GetNowTime() (Year, Month, Day, Hour, Minute, Second int) {
	Year = time.Now().Year()
	Month = int(time.Now().Month())
	Day = time.Now().Day()
	Hour = time.Now().Hour()
	Minute = time.Now().Minute()
	Second = time.Now().Second()
	return
}

//GetNowTimeInt8 设置20180611时间格式
func GetNowTimeInt8() int {
	return time.Now().Year()*10000 + int(time.Now().Month())*100 + time.Now().Day()
}

//GetNowTimeInt 设置2018061120时间格式
func GetNowTimeInt() int {
	return time.Now().Year()*1000000 + int(time.Now().Month())*10000 + time.Now().Day()*100 + time.Now().Hour()
}

//GetNowTimeInt64 设置20180626145950时间格式
func GetNowTimeInt64() int64 {
	y, m, d, h, mm, s := GetNowTime()
	return int64(s) + int64(mm)*100 + int64(h)*10000 + int64(d)*1000000 + int64(m)*100000000 + int64(y)*10000000000
}

//GetNowTimeString 设置时间格式为：2006-01-02 15:04:05
func GetNowTimeString() (tm string) {
	tm = time.Now().Format("2006-01-02 15:04:05")
	return
}

//IsLeapYear 判断是否为闰年
func IsLeapYear(year int) bool { //y == 2000, 2004
	//判断是否为闰年
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}
	return false
}

//GetDays 获取指定年月的天数
func GetDays(y, m int) (ds int) {
	ds = 31
	if m == 2 {
		if y%4 == 0 && y%100 != 0 || y%400 == 0 {
			ds = 29
		} else {
			ds = 28
		}
	} else if m == 4 || m == 6 || m == 9 || m == 11 {
		ds = 30
	}
	return
}

//TimeLoop 每隔h小时m分钟s秒就会执行一次f函数
func TimeLoop(f func(map[string]interface{}), h, m, s int, val map[string]interface{}) {
	t := time.NewTicker(time.Duration(h)*time.Hour + time.Duration(m)*time.Minute + time.Duration(s)*time.Second)
	for {
		select {
		case <-t.C:
			f(val)
		}
	}
}

//GetTime 返回指定时间的整型的年、月、日、时、分、秒
func GetTime(tm time.Time) (Year, Month, Day, Hour, Minute, Second int) {
	Year = tm.Year()
	Month = int(tm.Month())
	Day = tm.Day()
	Hour = tm.Hour()
	Minute = tm.Minute()
	Second = tm.Second()
	return
}
