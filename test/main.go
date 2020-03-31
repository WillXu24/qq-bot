package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

/*func main() {
	timeStr := time.Now().Format("2006-01-02")
	fmt.Println("timeStr:", timeStr)
	t, _ := time.Parse("2006-01-02", timeStr)
	timeNumber := t.Unix()
	fmt.Println("timeNumber:", timeNumber)
}
*/

/*func main() {
	t := time.Now()
	newTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	fmt.Println(newTime.Unix())
}*/

//func main() {
//	c := cron.New()
//
//	c.AddFunc("*/1 * * * * *", func() {
//		fmt.Println("every 1 seconds executing")
//	})
//
//	go c.Start()
//	defer c.Stop()
//
//	select {
//	case <-time.After(time.Second * 10):
//		return
//	}
//}

//定义函数类型

type Msg func(name string) string

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := &sync.WaitGroup{}
	c := make(chan os.Signal, 1)
	handleMap := make(map[int]Msg)
	handleMap[1] = handle1
	handleMap[2] = handle2
	handleMap[3] = handle3

	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		sig := <-c
		_ = sig
		s := handleMap[3]
		s("测试")
		wg.Done()
	}()
	wg.Add(1)
	fmt.Println("执行任务～～～")
	wg.Wait()
	fmt.Printf("结束")
}

func handle1(name string) string {
	fmt.Println(name)
	return "handle1"

}
func handle2(name string) string {
	fmt.Println("handle2")
	return "handle2"

}
func handle3(tt string) string {
	fmt.Println(tt)
	return "handle3"

}
