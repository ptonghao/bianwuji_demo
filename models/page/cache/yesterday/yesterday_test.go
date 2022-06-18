/*
 * @Author: Jimpu
 * @Description: test
 */

package yesterday

import (
	"bianwuji_demo/library/utils"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestYesterday(t *testing.T) {
	// 查询数据
	queryStart := time.Now()
	yes0 := utils.GetYesterday(0, 1, 2, 3)
	yes23 := utils.GetYesterday(23, 1, 2, 3)
	datas := GetTemperature(yes0, yes23)
	queryEnd := time.Since(queryStart)

	log.Println(fmt.Sprintf("TestYesterday 查询花费时间=%v", queryEnd))
	assert.True(t, len(datas) > 0, "test success!")

}
