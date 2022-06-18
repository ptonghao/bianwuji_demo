/*
 * @Author: Jimpu
 * @Description: debug data
 */

package cache

import (
	"bianwuji_demo/library/utils"
	"fmt"
	"math/rand"
	"time"
)

// init .
func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRand(num int) int {
	v := rand.Intn(num)
	return v
}

// MockCurrentData mock当前的数据
func MockCurrentData(pType string, callback func(string, int64, string) bool) {
	// 模拟当前的数据
	callback(pType, int64(getRand(100)), fmt.Sprintf("%v", getRand(100)))
}

// MockYesterdayData mock昨天的数据
func MockYesterdayData(pType string, callback func(string, int64, string) bool) {
	now := time.Now()
	yesDate := now.AddDate(0, 0, -1)
	yes0 := time.Date(yesDate.Year(), yesDate.Month(), yesDate.Day(), 0, 0, 0, 0, yesDate.Location())
	timeSamp := int64(yes0.Unix())
	// log.Println(fmt.Sprintf("MockYesterdayData len=%v", 3600*24))
	// 模拟一天的数据,假设每秒产生一条数据,每天一共有3600 * 24 条数据
	for i := 0; i < 3600*24; i++ {
		if !callback(pType, timeSamp+int64(i), fmt.Sprintf("%v", getRand(100))) {
			return
		}
	}
}

// MockLastMonthData mock 上个月数据
func MockLastMonthData(pType string, callback func(string, int64, string) bool) {
	lastMonthStart, lastMonthEnd := utils.GetLastMonth()
	// 模拟一个月的数据量, 一共 3600 * 24 * 30条
	// log.Println(fmt.Sprintf("MockLastMonthData lastMonthStart=%v, lastMonthEnd=%v, diff=%v", lastMonthStart, lastMonthEnd, lastMonthEnd-lastMonthStart))
	for i := lastMonthStart; i < lastMonthEnd; i++ {
		if !callback(pType, i, fmt.Sprintf("%v", getRand(100))) {
			return
		}
	}
}
