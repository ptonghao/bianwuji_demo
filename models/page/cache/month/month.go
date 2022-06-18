/*
 * @Author: Jimpu
 * @Description: 上个月的数据写入redis或者go-cache
 */

package month

import (
	"bianwuji_demo/models/page/cache"
	"bianwuji_demo/models/page/consts"
	"fmt"
	"log"
	"time"
)

// init .
func init() {
	// 当实例启动的时候,从db里取出yesterday的数据初始化内存
	LoadDatas()
}

var Temperature MonthData // 温度计温度
var Position MonthData    // 机械臂位置
var PlateReader MonthData // 酶标仪

// LoadDatas 从db加载上个月的数据
func LoadDatas() {
	// 加载温度计温度
	tStart := time.Now()
	LoadTemperature()
	log.Println(fmt.Sprintf("LoadDatas 加载上个月 [温度计温度] 数据花费时间=%v", time.Since(tStart)))

	// 加载机械臂位置
	pStart := time.Now()
	LoadPosition()
	log.Println(fmt.Sprintf("LoadDatas 加载上个月 [机械臂位置] 数据花费时间=%v", time.Since(pStart)))

	// 加载酶标仪
	prStart := time.Now()
	LoadPlateReader()
	log.Println(fmt.Sprintf("LoadDatas 加载上个月 [酶标仪] 数据花费时间=%v", time.Since(prStart)))
}

// GetTemperature 获取当前温度
func GetTemperature(startTime, endTime int64) []string {
	return Temperature.QueryDatas(startTime, endTime)
}

// GetPosition 获取当前机械臂位置
func GetPosition(startTime, endTime int64) []string {
	return Position.QueryDatas(startTime, endTime)
}

// GetPlateReader 获取当前酶标仪
func GetPlateReader(startTime, endTime int64) []string {
	return PlateReader.QueryDatas(startTime, endTime)
}

// 清理重更新数据
func CleanAndUpdateData() {
	// 清空内存数据
	Temperature.Clear()
	Position.Clear()
	PlateReader.Clear()

	// 从db加载数据到内存
	LoadTemperature()
	LoadPosition()
	LoadPlateReader()
}

// 从db加载数据到内存
func LoadTemperature() {
	// todo ...
	cache.MockLastMonthData(consts.P_TEMPERATURE, MockMonthDataCallBack)
}

// 从db加载数据到内存
func LoadPosition() {
	// todo ...
	cache.MockLastMonthData(consts.P_POSITION, MockMonthDataCallBack)
}

// 从db加载数据到内存
func LoadPlateReader() {
	// todo ...
	cache.MockLastMonthData(consts.P_PLATEREADER, MockMonthDataCallBack)
}

// mock数据
func MockMonthDataCallBack(pType string, timeSamp int64, value string) bool {
	if pType == consts.P_TEMPERATURE {
		Temperature.Add(timeSamp, value)
	} else if pType == consts.P_POSITION {
		Position.Add(timeSamp, value)
	} else if pType == consts.P_PLATEREADER {
		PlateReader.Add(timeSamp, value)
	}
	return true
}
