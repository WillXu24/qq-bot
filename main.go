package main

import (
	"github.com/gin-gonic/gin"
	"qq_bot/api"
)

func main() {
	// 路由部分
	r := gin.Default()
	r.POST("/qq", api.MsgHandler)
	r.Run()
}
