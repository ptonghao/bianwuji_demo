/*
 * @Author: Jimpu
 * @Description: test
 */

package month

import (
	"bianwuji_demo/library/utils"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMonth(t *testing.T) {
	// 查询数据
	queryStart := time.Now()
	lastMonthStart, lastMonthEnd := utils.GetLastMonth()
	datas := GetTemperature(lastMonthStart+1000, lastMonthEnd-1000)
	queryEnd := time.Since(queryStart)

	log.Println(fmt.Sprintf("TestMonth 查询总数据花费时间=%v[建议分页拉取], 数据条数为:%v", queryEnd, len(datas)))
	assert.True(t, len(datas) > 0, "test success!")
}
