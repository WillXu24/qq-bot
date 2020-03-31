package main

import (
	"github.com/gin-gonic/gin"
	"qq_bot/api"
)

func main() {
	// 消息路由
	r := gin.Default()
	r.POST("/qq", api.MsgHandler)
	r.Run()
}
