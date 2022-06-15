/*
 * @Author: Jimpu
 * @Description: 昨天数据存入内存
 */

package yesterday

import (
	"bianwuji_demo/library/utils"
	"bianwuji_demo/models/page/cache"
	"time"
)

// init .
func init() {
	// 当实例启动的时候,从db里取出yesterday的数据初始化内存
	LoadDatas()
}

// 温度计温度
var Temperature = YesterdayData{
	IdxDatas: make(map[int64]string),
}

// 机械臂位置
var Position = YesterdayData{
	IdxDatas: make(map[int64]string),
}

// 酶标仪
var PlateReader = YesterdayData{
	IdxDatas: make(map[int64]string),
}

// LoadDatas 从db加载昨天的数据
func LoadDatas() {
	// 加载温度计温度
	go func() {
		time.Sleep(1 * time.Second)
		LoadTemperature()
	}()

	// 加载机械臂位置
	go func() {
		time.Sleep(1 * time.Second)
		LoadPosition()
	}()

	// 加载酶标仪
	go func() {
		time.Sleep(1 * time.Second)
		LoadPlateReader()
	}()

	// 每天0点执行定时器任务,清理并重更新数据
	utils.ScheduleTask(time.Now().Day(), 0, 0, 0, CleanAndUpdateData)
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

// mock数据
func MockDataCallBack(pType string, timeSamp int64, value string) bool {
	if pType == "t" {
		Temperature.Add(timeSamp, value)
	} else if pType == "p" {
		Position.Add(timeSamp, value)
	} else if pType == "pr" {
		PlateReader.Add(timeSamp, value)
	}
	return true
}

// 从db加载数据到内存
func LoadTemperature() {
	// todo ...
	cache.MockData("t", MockDataCallBack)
}

// 从db加载数据到内存
func LoadPosition() {
	// todo ...
	cache.MockData("p", MockDataCallBack)
}

// 从db加载数据到内存
func LoadPlateReader() {
	// todo ...
	cache.MockData("pr", MockDataCallBack)
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
