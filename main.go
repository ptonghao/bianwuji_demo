/*
 * @Author: Jimpu
 * @Description: main
 */

package main

import (
	"bianwuji_demo/bootstrap"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default() // 创建引擎
	engine.POST("/query", bootstrap.QueryData)

	// 启动引擎，监听8888端口
	if err := engine.Run(":8888"); err != nil {
		log.Print(err)
	}
}
