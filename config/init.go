package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
	"log"
	"os"
	"time"
)

var DatabaseURL string
var DatabaseName string
var QQ string
var CoolQURL string

func init() {
	err := godotenv.Load("./config/config.env")
	if err != nil {
		log.Fatal("Config init failed:", err)
	}
	// 慢启动处理
	slowStart()

	// 初始化参数
	DatabaseURL = os.Getenv("DB_URI")
	DatabaseName = os.Getenv("DB_NAME")
	QQ = os.Getenv("QQ")
	CoolQURL = os.Getenv("COOLQ_URL")
}

func slowStart() {
	c := cron.New()

	c.AddFunc("*/10 * * * * *", func() {
		fmt.Println("启动准备中")
	})

	go c.Start()
	defer c.Stop()

	select {
	case <-time.After(time.Second * 50):
		return
	}
}
