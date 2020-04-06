package logs

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func init() {
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
	return
}
