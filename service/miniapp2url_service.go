package service

import (
	"fmt"
	"log"
	"regexp"
)

var titleRegex = "\"title\":\"(.*?)\""
var descRegex = "\"desc\":\"(.*?)\""
var urlRegex = "\"qqdocurl\":\"(.*?)\""


// 小程序转链接
func MiniApp2URL(groupID int64, msg string) {
	// 匹配标题
	reg := regexp.MustCompile(titleRegex)
	title:= reg.FindStringSubmatch(msg)
	if title==nil{
		log.Println("未匹配到标题")
		return
	}
	// 匹配描述
	reg = regexp.MustCompile(descRegex)
	desc := reg.FindStringSubmatch(msg)
	if desc==nil{
		log.Println("未匹配到描述")
		return
	}
	// 匹配链接
	reg = regexp.MustCompile(urlRegex)
	url := reg.FindStringSubmatch(msg)
	if url==nil{
		log.Println("未匹配到链接")
		//url=append(url,"","由于腾讯在安卓系统做了限制，暂时无法获取链接，换个苹果再来吧")
		return
	}
	// 生成回复
	res := fmt.Sprintf("【小程序转链接】\n【来自】：%s\n\n【标题】：%s\n\n【链接】：%s\n\n",title[1],desc[1],url[1])
	Send2group(groupID,res)
}