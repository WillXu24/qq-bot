package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
	"log"
	"os"
	"strconv"
	"time"
)

var DatabaseURL string
var DatabaseName string
var CoolQURL string

var BotQQ int64
var AdminQQ int64

func init() {
	// 读取配置文件
	err := godotenv.Load("./config/config.env")
	if err != nil {
		log.Fatal("Config init failed(1):", err)
	}
	// 慢启动处理
	slowStart()

	// 初始化字符串参数
	DatabaseURL = os.Getenv("DB_URI")
	DatabaseName = os.Getenv("DB_NAME")
	CoolQURL = os.Getenv("COOLQ_URL")

	// 初始化整形参数
	botQQ, err := strconv.ParseInt(os.Getenv("BOT_QQ"), 10, 64)
	if err != nil {
		log.Fatal("Config init failed(2):", err)
	}
	adminQQ, err := strconv.ParseInt(os.Getenv("BOT_QQ"), 10, 64)
	if err != nil {
		log.Fatal("Config init failed(3):", err)
	}
	BotQQ = botQQ
	AdminQQ = adminQQ
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
