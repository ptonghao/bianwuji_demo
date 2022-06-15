/*
 * @Author: Jimpu
 * @Description: 产品工厂
 */

package products

import (
	"bianwuji_demo/models/page/consts"
	"bianwuji_demo/models/page/service/products/plate_reader"
	"bianwuji_demo/models/page/service/products/position"
	"bianwuji_demo/models/page/service/products/temperature"
	"errors"
	"fmt"
)

// IProducts 查询产品接口
type IProducts interface {
	Query(dim string, begin, end int64) []string
}

// ProductMaps 产品集合
var ProductMaps = map[string]IProducts{
	consts.P_TEMPERATURE: temperature.NewTemperature(),
	consts.P_POSITION:    position.NewPosition(),
	consts.P_PLATEREADER: plate_reader.NewPlateReader(),
}

// Products 查询产品参数
type Products struct {
	Product   string
	Dim       string
	BeginTime int64
	EndTime   int64
}

// QueryProducts 查询产品
func (p *Products) QueryProducts() (ret []string, err error) {
	v, ok := ProductMaps[p.Product]
	if !ok {
		err = errors.New("not found product")
		return
	}
	ret = v.Query(p.Dim, p.BeginTime, p.EndTime)

	// todo .. 如果查询没有命中缓存,应该去查询db(降级兜底方案)
	if len(ret) == 0 {
		err = fmt.Errorf("查询没有命中缓存,应该去查询db(降级兜底方案)")
	}
	return
}
