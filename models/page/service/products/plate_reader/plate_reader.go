/*
 * @Author: Jimpu
 * @Description: 酶标仪
 */

package plate_reader

import (
	current_cache "bianwuji_demo/models/page/cache/current"
	month_cache "bianwuji_demo/models/page/cache/month"
	day_cache "bianwuji_demo/models/page/cache/yesterday"
	"bianwuji_demo/models/page/consts"
)

type PlateReader struct {
}

func NewPlateReader() *PlateReader {
	return &PlateReader{}
}

func (c *PlateReader) Query(dim string, begin, end int64) (ret []string) {
	if dim == consts.DIM_MONTH {
		// 查询月数据
		ret = month_cache.GetPlateReader(begin, end)
	} else if dim == consts.DIM_DAY {
		// 查询天数据
		ret = day_cache.GetPlateReader(begin, end)
	} else if dim == consts.DIM_CURRENT {
		// 查询当前数据
		ret = current_cache.GetPlateReader()
	}

	return
}
