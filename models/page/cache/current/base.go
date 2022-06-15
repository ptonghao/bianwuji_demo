/*
 * @Author: Jimpu
 * @Description: 缓存当前数据
 */

package current

import "sync"

// CurrentData 存当前值结构体
type CurrentData struct {
	Value string
	Lock  sync.RWMutex
}

// Set .
func (c *CurrentData) Set(value string) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Value = value
}

// Get .
func (c *CurrentData) Get() string {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	return c.Value
}
