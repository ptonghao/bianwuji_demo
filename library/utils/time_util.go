/*
 * @Author: Jimpu
 * @Description: time util
 */

package utils

import "time"

// SetTime 获取本月某一天的点时间
func SetTime(day, hour, min, second int) (d time.Duration) {
	now := time.Now()
	setTime := time.Date(now.Year(), now.Month(), day, hour, min, second, 0, now.Location())
	d = setTime.Sub(now)
	if d > 0 {
		return
	}
	return d + time.Hour*24
}

// GetYesterday 获取昨天某个时间点的时间
func GetYesterday(hour, min, sec, msec int) int64 {
	now := time.Now()
	yesDate := now.AddDate(0, 0, -1)
	yes0 := time.Date(yesDate.Year(), yesDate.Month(), yesDate.Day(), hour, min, sec, msec, yesDate.Location())
	return int64(yes0.Unix())
}

// 获取上一个月的开始 结束 时间戳
func GetLastMonth() (int64, int64) {
	now := time.Now()
	lastMonthFirstDay := now.AddDate(0, -1, -now.Day()+1)
	lastMonthStart := time.Date(lastMonthFirstDay.Year(), lastMonthFirstDay.Month(), lastMonthFirstDay.Day(), 0, 0, 0, 0, now.Location()).Unix()
	lastMonthEndDay := lastMonthFirstDay.AddDate(0, 1, -1)
	lastMonthEnd := time.Date(lastMonthEndDay.Year(), lastMonthEndDay.Month(), lastMonthEndDay.Day(), 23, 59, 59, 0, now.Location()).Unix()
	return lastMonthStart, lastMonthEnd
}
