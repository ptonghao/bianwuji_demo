/*
 * @Author: Jimpu
 * @Description: test
 */

package yesterday

import (
	"bianwuji_demo/library/utils"
	"testing"
)

func init() {
	LoadDatas()
}

func BenchmarkYesterday(b *testing.B) {
	b.ResetTimer()

	yes0 := utils.GetYesterday(0, 1, 2, 3)
	yes23 := utils.GetYesterday(23, 1, 2, 3)
	for i := 0; i < b.N; i++ {
		GetTemperature(yes0, yes23)
	}
	b.StopTimer()
}
