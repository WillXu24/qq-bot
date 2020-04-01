package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"qq_bot/api"
	"qq_bot/service"
)

func main() {
	// 定时任务
	timedTasks()
	// 消息路由
	r := gin.Default()
	r.POST("/qq", api.MsgHandler)
	r.Run()
}

func timedTasks() {
	c := cron.New()

	// 每日龙王推送
	c.AddFunc("0 59 23 * * *", service.CounterClear)

	go c.Start()
	defer c.Stop()
}
