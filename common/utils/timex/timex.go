package timex

import "time"

// GetDisTodayEnd 获取现在到今天结束的秒数
func GetDisTodayEnd() int64 {
	todayEnd, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	todayEndUnix := todayEnd.AddDate(0, 0, 1).Unix()
	period := todayEndUnix - time.Now().Unix()
	return period
}
