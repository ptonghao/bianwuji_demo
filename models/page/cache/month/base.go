/*
 * @Author: Jimpu
 * @Description: go-cache或者redis
 */

package month

import (
	"bianwuji_demo/library/utils"
	"fmt"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

// MonthData 存储当天数据,有序的
type MonthData struct {
	Lock      sync.RWMutex
	TimeSlice []int64 // 顺序存储上个月的数据记录的时间戳

	// go-cache 本地缓存
	Datas *cache.Cache
	Once  sync.Once
}

// Add 插入数据
func (c *MonthData) Add(timestamp int64, value string) {
	c.Once.Do(func() {
		// 默认过期时间为1day，每5min清理一次过期缓存
		c.Datas = cache.New(1*time.Hour*24, 5*time.Minute)
	})

	c.Lock.Lock()
	defer c.Lock.Unlock()

	c.TimeSlice = append(c.TimeSlice, timestamp)
	c.Datas.Set(fmt.Sprintf("%v", len(c.TimeSlice)-1), value, 1*time.Hour*24)
}

// QueryDatas 查询昨天的数据
func (c *MonthData) QueryDatas(startTime, endTime int64) (ret []string) {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	if len(c.TimeSlice) == 0 {
		return
	}

	// start := time.Now()
	minIdx := utils.FindBinarySearch(c.TimeSlice, 0, int64(len(c.TimeSlice)-1), utils.Max(c.TimeSlice[0], startTime))
	maxIdx := utils.FindBinarySearch(c.TimeSlice, 0, int64(len(c.TimeSlice)-1), utils.Min(c.TimeSlice[len(c.TimeSlice)-1], endTime))
	// log.Println(fmt.Sprintf("QueryDatas TimeSlice.len=%v, minIdx=%v, maxIdx=%v, startTime=%v, endTime=%v, time=%v", len(c.TimeSlice), minIdx, maxIdx, startTime, endTime, time.Since(start)))
	if minIdx == -1 || maxIdx == -1 {
		return
	}
	if startTime == endTime {
		if v, ok := c.Datas.Get(fmt.Sprintf("%v", minIdx)); ok {
			ret = append(ret, fmt.Sprintf("%v", v))
		}
		return
	}
	for idx := minIdx; idx <= maxIdx; idx++ {
		if v, ok := c.Datas.Get(fmt.Sprintf("%v", idx)); ok {
			ret = append(ret, fmt.Sprintf("%v", v))
		}
	}
	return
}

// Clear 清空内存数据
func (c *MonthData) Clear() {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	c.TimeSlice = []int64{}
}

// var RedisCache = &redis.Client{}

// func init() {
// 	RedisCache = redis.NewClient(&redis.Options{
// 		Addr:     "127.0.0.1:6379",
// 		Password: "123456",
// 		DB:       0,
// 	})

// 	//ping
// 	pong, err := RedisCache.Ping().Result()
// 	if err != nil {
// 		fmt.Println("ping error", err.Error())
// 		return
// 	}
// 	fmt.Println("ping result:", pong)
// }

// // Get 获取缓存数据
// func Get(key string) (string, error) {
// 	result, err := RedisCache.Get(key).Result()
// 	return result, err
// }

// // Set 设置数据 过期时间默认24H
// func Set(key string, value interface{}) error {
// 	err := RedisCache.Set(key, value, time.Hour*24).Err()
// 	return err
// }

// // LPush RPush 使用RPush命令往队列右边加入
// func LPush(key string, value ...interface{}) error {
// 	err := RedisCache.LPush(key, value).Err()
// 	return err
// }

// // RPop LPop 取出并移除左边第一个元素
// func RPop(key string) (interface{}, error) {
// 	result, err := RedisCache.RPop(key).Result()
// 	return result, err
// }

// // BRPop BLPop 取出并移除左边第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
// func BRPop(timeout time.Duration, key string) (interface{}, error) {
// 	result, err := RedisCache.BRPop(timeout, key).Result()
// 	return result, err
// }

// // LLen 获取数据长度
// func LLen(key string) (int64, error) {
// 	result, err := RedisCache.LLen(key).Result()
// 	return result, err
// }

// // LRange 获取数据列表
// func LRange(key string, start, end int64) ([]string, error) {
// 	result, err := RedisCache.LRange(key, start, end).Result()
// 	return result, err
// }

// // HSet hash相关操作
// //set hash 适合存储结构
// func HSet(hashKey, key string, value interface{}) error {
// 	err := RedisCache.HSet(hashKey, key, value).Err()
// 	return err
// }

// // HGet get Hash
// func HGet(hashKey, key string) (interface{}, error) {
// 	result, err := RedisCache.HGet(hashKey, key).Result()
// 	return result, err
// }

// // HGetAll 获取所以hash ,返回map
// func HGetAll(hashKey string) (map[string]string, error) {
// 	result, err := RedisCache.HGetAll(hashKey).Result()
// 	return result, err
// }

// // HDel 删除一个或多个哈希表字段
// func HDel(hashKey string, key ...string) error {
// 	err := RedisCache.HDel(hashKey, key...).Err()
// 	return err
// }

// // HExists 查看哈希表的指定字段是否存在
// func HExists(hashKey, key string) (bool, error) {
// 	result, err := RedisCache.HExists(hashKey, key).Result()
// 	return result, err
// }

// // SAdd -----------------Set------------------------
// // 添加Set
// func SAdd(key string, values ...interface{}) error {
// 	err := RedisCache.SAdd(key, values).Err()
// 	return err
// }

// // SCard 获取集合的成员数
// func SCard(key string) (int64, error) {
// 	result, err := RedisCache.SCard(key).Result()
// 	return result, err
// }

// // SMembers 获取集合的所有成员
// func SMembers(key string) ([]string, error) {
// 	result, err := RedisCache.SMembers(key).Result()
// 	return result, err
// }

// // SRem 移除集合里的某个元素
// func SRem(key string, value interface{}) error {
// 	err := RedisCache.SRem(key, value).Err()
// 	return err
// }

// // SPop 移除并返回set的一个随机元素(SET是无序的)
// func SPop(key string) (interface{}, error) {
// 	result, err := RedisCache.SPop(key).Result()
// 	return result, err
// }

// // ZAdd ------------------ZSet-------------------------
// func ZAdd(key string, values []redis.Z) error {
// 	err := RedisCache.ZAdd(key, values...).Err()
// 	return err
// }

// // ZIncrBy 给指定的key和值加分
// func ZIncrBy(key string, score float64, value string) error {
// 	err := RedisCache.ZIncrBy(key, score, value).Err()
// 	return err
// }

// // ZRevRangeWithScores 取zSet里的前n名热度的数据
// func ZRevRangeWithScores(key string, start, end int64) ([]redis.Z, error) {
// 	result, err := RedisCache.ZRevRangeWithScores(key, start, end).Result()
// 	return result, err
// }

// // Expire 给指定key 设置过期时间
// func Expire(key string, duration time.Duration) error {
// 	err := RedisCache.Expire(key, duration).Err()
// 	return err
// }

// // ExpireAt 给指定Key 设置过期时间，时间格式为UNIX时间
// func ExpireAt(key string, duration time.Time) error {
// 	err := RedisCache.ExpireAt(key, duration).Err()
// 	return err
// }

// // TTL 获取key的生存时间
// func TTL(key string) (time.Duration, error) {
// 	result, err := RedisCache.TTL(key).Result()
// 	return result, err
// }
