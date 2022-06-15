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

// ScheduleTask 新建定时器
func ScheduleTask(day, hour, min, second int, callback func()) {
	go func() {
		// 每天0时0分触发更新
		t := time.NewTimer(SetTime(day, hour, min, second))
		defer t.Stop()
		for {
			select {
			case <-t.C:
				// t.Reset(time.Hour * 24)
				// 定时任务函数
				callback()
			}
		}
	}()
}
