package config

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
	"io"
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
	// 初始化日志文件
	logInit()

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

func logInit() {
	// 日志初始化
	file := "./logs/access.log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.SetPrefix("[log] ")
	//log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	// 路由访问日志
	//gin.DisableConsoleColor()
	//file, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
}

func slowStart() {
	c := cron.New()

	c.AddFunc("*/10 * * * * *", func() {
		log.Println("启动准备中")
	})

	go c.Start()
	defer c.Stop()

	select {
	case <-time.After(time.Second * 50):
		return
	}
}
