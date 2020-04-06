package service

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"qq_bot/config"
)

func Send2person(id int64, msg string) {
	// 初始化请求
	client := &http.Client{}
	// 初始化json
	sendMsg := make(map[string]interface{})
	sendMsg["user_id"] = id
	sendMsg["message"] = msg
	marshal, err := json.Marshal(sendMsg)
	if err != nil {
		log.Println(err)
	}
	// 提交请求
	request, err := http.NewRequest("POST", config.CoolQURL+"/send_private_msg", bytes.NewReader(marshal))
	if err != nil {
		log.Println(err)
	}
	request.Header.Add("Content-Type", "application/json")
	// 处理返回结果
	_, err = client.Do(request)
	if err != nil {
		log.Println(err)
	}
}

func Send2group(id int64, msg string) {
	// 初始化请求
	client := &http.Client{}
	// 初始化json
	sendMsg := make(map[string]interface{})
	sendMsg["group_id"] = id
	sendMsg["message"] = msg
	marshal, err := json.Marshal(sendMsg)
	if err != nil {
		log.Println(err)
	}
	// 提交请求
	request, err := http.NewRequest("POST", config.CoolQURL+"/send_group_msg", bytes.NewReader(marshal))
	if err != nil {
		log.Println(err)
	}
	request.Header.Add("Content-Type", "application/json")
	// 处理返回结果
	_, err = client.Do(request)
	if err != nil {
		log.Println(err)
	}
}
