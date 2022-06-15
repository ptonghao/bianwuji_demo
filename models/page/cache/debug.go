/*
 * @Author: Jimpu
 * @Description: debug data
 */

package cache

import (
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

func MockData(pType string, callback func(string, int64, string) bool) {
	now := time.Now()
	yesDate := now.AddDate(0, 0, -1)
	yes0 := time.Date(yesDate.Year(), yesDate.Month(), yesDate.Day(), 0, 0, 0, 0, yesDate.Location())
	timeSamp := int64(yes0.Unix())
	// fmt.Println(fmt.Sprintf("MockData timeSamp=%v", timeSamp))
	for i := 0; i < 3600*24; i++ {
		if !callback(pType, timeSamp+int64(i), fmt.Sprintf("%v", getRand(100))) {
			return
		}
	}
}
