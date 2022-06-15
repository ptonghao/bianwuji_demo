/*
 * @Author: Jimpu
 * @Description: 存放昨天的数据
 */

package yesterday

import (
	"bianwuji_demo/library/utils"
	"sync"
)

// YesterdayData 存储当天数据,有序的
type YesterdayData struct {
	TimeSlice []int64          // 顺序存储昨天的数据记录的时间戳
	IdxDatas  map[int64]string // 存储slice下标, key: 下标, value: 采集的数据
	Lock      sync.RWMutex
}

// Add 插入数据
func (c *YesterdayData) Add(timestamp int64, value string) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	c.TimeSlice = append(c.TimeSlice, timestamp)
	c.IdxDatas[int64(len(c.TimeSlice)-1)] = value
}

// QueryDatas 查询昨天的数据
func (c *YesterdayData) QueryDatas(startTime, endTime int64) (ret []string) {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	if len(c.TimeSlice) == 0 {
		return
	}

	minIdx := utils.FindBinarySearch(c.TimeSlice, 0, int64(len(c.TimeSlice)-1), utils.Max(c.TimeSlice[0], startTime))
	maxIdx := utils.FindBinarySearch(c.TimeSlice, 0, int64(len(c.TimeSlice)-1), utils.Min(c.TimeSlice[len(c.TimeSlice)-1], endTime))
	if minIdx == -1 || maxIdx == -1 {
		return
	}
	if startTime == endTime {
		ret = append(ret, c.IdxDatas[int64(minIdx)])
		return
	}

	for v := minIdx; v <= maxIdx; v++ {
		ret = append(ret, c.IdxDatas[int64(v)])
	}
	return
}

// Clear 清空内存数据
func (c *YesterdayData) Clear() {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	c.TimeSlice = []int64{}
	for k, _ := range c.IdxDatas {
		delete(c.IdxDatas, k)
	}
}
