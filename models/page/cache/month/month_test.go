/*
 * @Author: Jimpu
 * @Description: test
 */

package month

import (
	"bianwuji_demo/library/utils"
	"testing"
)

func init() {
	LoadDatas()
}

func BenchmarkMonth(b *testing.B) {
	b.ResetTimer()
	lastMonthStart, lastMonthEnd := utils.GetLastMonth()
	for i := 0; i < b.N; i++ {
		GetTemperature(lastMonthStart, lastMonthEnd)
	}
	b.StopTimer()
}
