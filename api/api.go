package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"qq_bot/service"
	"qq_bot/views"
	"strings"
)

func MsgHandler(c *gin.Context) {
	// 信息绑定
	var msg views.PostMsg
	err := c.ShouldBind(&msg)
	if err != nil {
		fmt.Println("ShouldBind:", err)
	}

	// 测试用
	fmt.Println(msg.SubType)
	fmt.Println(msg.Message)
	fmt.Println(msg.RawMessage)
	fmt.Println(msg.Sender)

	// 群聊
	if msg.MessageType == "group" {
		// 基本功能：消息计数
		err := service.MsgCounter(msg)
		if err != nil {
			fmt.Println("MsgCounter", err)
		}
		// 附加功能：
		res := funcSelector(msg.RawMessage)
		if res == "今日龙王" {
			err := service.Send2group(msg.GroupId, "我明白了,你想查看"+res+"对不对")
			if err != nil {
				fmt.Println("Send2group", err)
			}
			err = service.GetDragonToday(msg.GroupId)
			if err != nil {
				fmt.Println("GetDragonToday", err)
			}
			return
		}
		if res == "历史龙王" {
			err := service.Send2group(msg.GroupId, "我明白了,你想查看"+res+"对不对")
			if err != nil {
				fmt.Println("Send2group", err)
			}
			err = service.GetDragonHistory(msg.GroupId)
			if err != nil {
				fmt.Println("GetDragonHistory", err)
			}
			return
		}
		// 默认功能
	}

	// 私聊
	if msg.MessageType == "private" {
		// 基本功能：复读机
		err := service.Send2person(msg.UserId, msg.RawMessage)
		if err != nil {
			fmt.Println("Send2person", err)
		}
	}
}

func funcSelector(msg string) string {
	var funcMap = map[string]int{
		"今日龙王": 0,
		"历史龙王": 0,
	}
	if index := strings.Index(msg, "龙王"); index > 0 {
		funcMap["今日龙王"]++
		funcMap["历史龙王"]++
	}
	if index := strings.Index(msg, "今"); index > 0 {
		funcMap["今日龙王"]++
	}
	if index := strings.Index(msg, "历史"); index > 0 {
		funcMap["历史龙王"]++
	}
	if index := strings.Index(msg, "老"); index > 0 {
		funcMap["历史龙王"]++
	}

	var res = "今日龙王"
	if funcMap["历史龙王"] > funcMap["今日龙王"] {
		res = "历史龙王"
	}
	return res
}
