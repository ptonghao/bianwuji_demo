/*
 * @Author: Jimpu
 * @Description: scheduler
 */

package bootstrap

import (
	"bianwuji_demo/models/page/cache/yesterday"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
)

type CronTab struct {
	Cron *cron.Cron
}

var once sync.Once

var instance *CronTab

// InitScheduler 任务调度器
func InitScheduler() *CronTab {
	instance = &CronTab{}
	once.Do(instance.createCron)
	return instance
}

func StopCron(cronTab *CronTab) {
	// 捕捉退出信号
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	<-c
	_ = cronTab.StopCron()
}

/**
 * @description: 初始化
 */
func (cronTab *CronTab) createCron() {
	option := cron.WithSeconds()
	chain := cron.WithChain(
		cron.Recover(cron.DefaultLogger), //use cron.DefaultLogger
	)
	cronTab.Cron = cron.New(chain, option)

	cronTab.ExecCronTask()

}

/*
 1）星号(*)
 表示 cron 表达式能匹配该字段的所有值。如在第5个字段使用星号(month)，表示每个月
 2）斜线(/)
 表示增长间隔，如第1个字段(minutes) 值是 3-59/15，表示每小时的第3分钟开始执行一次，之后每隔 15 分钟执行一次（即 3、18、33、48 这些时间点执行），这里也可以表示为：3/15
 3）逗号(,)
 用于枚举值，如第6个字段值是 MON,WED,FRI，表示 星期一、三、五 执行
 4）连字号(-)
 表示一个范围，如第3个字段的值为 9-17 表示 9am 到 5pm 直接每个小时（包括9和17）
 5）问号(?)
 只用于日(Day of month)和星期(Day of week)，\表示不指定值，可以用于代替 *

*/
// cron举例说明
//　　 每隔5秒执行一次：*/5 * * * * ?
//             每隔1分钟执行一次：0 */1 * * * ?
//             每天23点执行一次：0 0 23 * * ?
//             每天凌晨1点执行一次：0 0 1 * * ?
//             每月1号凌晨1点执行一次：0 0 1 1 * ?
//             在26分、29分、33分执行一次：0 26,29,33 * * * ?
//             每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?
func (cronTab *CronTab) ExecCronTask() {
	// 更新天级别定时任务,每天0点执行一次
	cronDay := "0 0 0 * * ?"
	dayTaskId, err := cronTab.Cron.AddFunc(cronDay, func() {
		start := time.Now()
		// 从db读取已排序的记录更新内存
		yesterday.CleanAndUpdateData()
		log.Println("[cron:%v] End cost:[%dms]", cronDay, time.Since(start)/1000000)
	})
	log.Println("Cron:%v: instance:%v err: %v", cronDay, dayTaskId, err)

	// 更新月级别定时任务,每月第一天0点执行一次
	cronMonth := "0 0 0 0 1 ?"
	monthTaskId, err := cronTab.Cron.AddFunc(cronDay, func() {
		start := time.Now()
		// todo ... 从db读取已排序的记录更新内存
		log.Println("[cron:%v] End cost:[%dms]", cronMonth, time.Since(start)/1000000)
	})
	log.Println("Cron:%v: instance:%v err: %v", cronMonth, monthTaskId, err)
}

func (cronTab *CronTab) StartCron() error {
	log.Println("Start Cron Task...")
	cronTab.Cron.Start()
	log.Println("Start Cron Task...done")
	return nil
}
func (cronTab *CronTab) StopCron() error {
	log.Println("Quit Cron Task...")
	cronTab.Cron.Stop()
	log.Println("Quit Cron Task... done")
	return nil
}
