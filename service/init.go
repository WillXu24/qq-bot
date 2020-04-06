package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"net/http"
	"qq_bot/config"
	"qq_bot/models"
)

// get_group_list
type groupRes struct {
	Data    []groupMsg `json:"data"`
	Retcode int64      `json:"retcode"`
	Status  string     `json:"status"`
}
type groupMsg struct {
	GroupID   int64  `json:"group_id"`
	GroupName string `json:"group_name"`
}

// get_group_member_list
type memberRes struct {
	Data    []memberMsg `json:"data"`
	Retcode int64       `json:"retcode"`
	Status  string      `json:"status"`
}
type memberMsg struct {
	GroupID  int64  `json:"group_id"`
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
}

// 初始化服务
func init() {
	GroupAndMemberInit()
}

func GroupAndMemberInit() {
	// 获取群组列表
	group, err := getGroup()
	if err != nil {
		log.Fatal("Service init failed:", err)
	}
	for i := range group {
		// 获取群成员列表
		member, err := getMember(group[i].GroupID)
		if err != nil {
			log.Fatal("Service init failed:", err)
		}
		// 查询是否初始化
		_, err = models.MemberFindOne(bson.M{"group_id": group[i].GroupID, "user_id": member[0].UserID})
		if err == nil {
			continue
		}
		// 不存在则初始化
		_, err = models.GroupInsertOne(models.GroupMsg{
			GroupId:   group[i].GroupID,
			GroupName: group[i].GroupName,
		})
		if err != nil {
			log.Fatal("Service init failed:", err)
		}
		for j := range member {
			// 跳过机器人qq
			if member[j].UserID == config.BotQQ {
				continue
			}
			// 写入数据库
			_, err := models.MemberInsertOne(models.MemberMsg{
				GroupId:   member[j].GroupID,
				GroupName: group[i].GroupName,
				UserId:    member[j].UserID,
				Username:  member[j].Nickname,
			})
			if err != nil {
				log.Fatal("Service init failed:", err)
			}
		}
	}
}

// 查找群组列表
func getGroup() ([]groupMsg, error) {
	// 初始化请求
	client := &http.Client{}
	// 提交请求
	request, err := http.NewRequest("GET", config.CoolQURL+"/get_group_list", nil)
	if err != nil {
		return nil, err
	}
	//request.Header.Add("Content-Type", "application/json")
	// 处理返回结果
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	// 解析json
	var res groupRes
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	if res.Retcode != 0 {
		return nil, errors.New("获取群组列表失败")
	}
	return res.Data, nil
}

// 查找群成员
func getMember(id int64) ([]memberMsg, error) {
	// 初始化请求
	client := &http.Client{}
	// 初始化json
	sendMsg := make(map[string]interface{})
	sendMsg["group_id"] = id
	marshal, err := json.Marshal(sendMsg)
	if err != nil {
		return nil, err
	}
	// 提交请求
	request, err := http.NewRequest("POST", config.CoolQURL+"/get_group_member_list", bytes.NewReader(marshal))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	// 处理返回结果
	response, _ := client.Do(request)
	body, err := ioutil.ReadAll(response.Body)
	// 解析json
	var res memberRes
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	if res.Retcode != 0 {
		return nil, errors.New("获取群成员信息失败")
	}
	return res.Data, nil
}
