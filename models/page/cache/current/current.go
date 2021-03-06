/*
 * @Author: Jimpu
 * @Description: 缓存当前数据
 */

package current

import (
	"bianwuji_demo/models/page/cache"
	"bianwuji_demo/models/page/consts"
)

var Temperature CurrentData // 温度计温度
var Position CurrentData    // 机械臂位置
var PlateReader CurrentData // 酶标仪

func MockDataCallBack(pType string, timeSamp int64, value string) bool {
	if pType == consts.P_TEMPERATURE {
		Temperature.Set(value)
	} else if pType == consts.P_POSITION {
		Position.Set(value)
	} else if pType == consts.P_PLATEREADER {
		PlateReader.Set(value)
	}
	return false
}

func init() {
	// 当实例启动的时候,从db里取出当前数据初始化
	// todo ...
	cache.MockCurrentData(consts.P_TEMPERATURE, MockDataCallBack)
	cache.MockCurrentData(consts.P_POSITION, MockDataCallBack)
	cache.MockCurrentData(consts.P_PLATEREADER, MockDataCallBack)
}

// SetTemperature 设置温度
func SetTemperature(data string) {
	Temperature.Set(data)
}

// SetPosition 设置机械臂位置
func SetPosition(data string) {
	Position.Set(data)
}

// SetPlateReader 设置酶标仪
func SetPlateReader(data string) {
	PlateReader.Set(data)
}

// GetTemperature 获取当前温度
func GetTemperature() []string {
	return []string{Temperature.Get()}
}

// GetPosition 获取当前机械臂位置
func GetPosition() []string {
	return []string{Position.Get()}
}

// GetPlateReader 获取当前酶标仪
func GetPlateReader() []string {
	return []string{PlateReader.Get()}
}
