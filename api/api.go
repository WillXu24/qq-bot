package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"qq_bot/service"
	"qq_bot/views"
	"strings"
)

func MsgHandler(c *gin.Context) {
	// 信息绑定
	var msg views.PostMsg
	err := c.ShouldBind(&msg)
	if err != nil {
		log.Println("ShouldBind:", err)
	}

	// 测试用
	log.Println("【消息类型】",msg.MessageType)
	log.Println("【消息子类型】",msg.SubType)
	log.Println("【消息内容】",msg.RawMessage)
	log.Println("【发送人信息】",msg.Sender)

	// 请求
	if msg.MessageType=="request"{
		service.GroupAndMemberInit()
	}

	// 群聊
	// 跳过匿名消息，或者采用事件过滤器：https://cqhttp.cc/docs/4.14/#/EventFilter
	if msg.MessageType == "group" && msg.SubType != "anonymous" {
		group(msg)
	}

	// 私聊
	if msg.MessageType == "private" {
		private(msg)
	}
}

func private(msg views.PostMsg) {
	// 基本功能：复读机
	service.Send2person(msg.UserId, msg.RawMessage)
}

func group(msg views.PostMsg) {
	// 基本功能：消息计数
	service.Counter(msg.GroupId, msg.UserId)
	// 附加功能：
	res := funcSelector(msg.RawMessage)
	if res == "今日龙王" {
		service.CounterRankToday(msg.GroupId)
		return
	}
	if res == "历史龙王" {
		service.CounterRankHistory(msg.GroupId)
		return
	}
	// 默认功能
}

func funcSelector(msg string) string {
	msg = " " + msg
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

	var res string
	for i, v := range funcMap {
		if v > 1 {
			res = i
		}
	}
	return res
}
