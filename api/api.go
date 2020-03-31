package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"qq_bot/service"
	"qq_bot/views"
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
		// 附加功能：查询龙王
		if msg.RawMessage == "今日龙王" {
			err = service.GetDragonToday(msg.GroupId)
			if err != nil {
				fmt.Println("MsgCounter", err)
			}
		}
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
