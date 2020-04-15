package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"qq_bot/models"
	"time"
)

// 消息计数
func Counter(groupID, userID int64) {
	// 初始化filter
	memberFilter := bson.M{"group_id": groupID, "user_id": userID}
	// 更新，返回
	_, err := models.MemberUpdateOne(memberFilter, bson.M{
		"$set": bson.M{
			"last_msg_at": time.Now().Unix()},
		"$inc": bson.M{
			"today_msg_count": 1,
			"total_msg_count": 1,
		}})
	if err != nil {
		log.Println("Counter:", err)
		return
	}
}

// 计数上报并清零
func CounterClear() {
	// 发送今日统计
	group, err := models.GroupFindMany(nil, nil)
	for i := range group {
		// 查询数据
		res, _ := models.MemberFindMany(bson.M{"group_id": group[i].GroupId, "today_msg_count": bson.M{"$gt": 0}}, &options.FindOptions{Sort: bson.M{"today_msg_count": -1}})
		// 没人发言或者发言少则不推送
		if res == nil || res[0].TodayMsgCount < 50 {
			continue
		}
		// 生成回复
		var msg string
		msg = fmt.Sprintf(
			"【每日龙王推送】\n%s今天说了%d句骚话，恭喜这个比！", res[0].Username, res[0].TodayMsgCount)
		// 调用发送函数
		Send2group(group[i].GroupId, msg)
	}
	// 清零所有成员的今日计数
	_, err = models.MemberUpdateMany(nil, bson.M{
		"$set": bson.M{
			"today_msg_count": 0,
		}})
	if err != nil {
		log.Println("Counter:", err)
	}
}

// 今日龙王
func CounterRankToday(groupId int64) {
	// 查询数据
	res, err := models.MemberFindMany(bson.M{"group_id": groupId, "today_msg_count": bson.M{"$gt": 0}}, &options.FindOptions{Sort: bson.M{"today_msg_count": -1}})
	if err != nil {
		log.Println("CounterRankToday(1):", err)
	}
	// 生成回复
	var msg string
	for i := range res {
		msg = msg + fmt.Sprintf("%d.%s:%d条\n", i+1, res[i].Username, res[i].TodayMsgCount)
	}
	msg = "今日骚话排行榜：\n" + msg + fmt.Sprintf("%s，给兄弟萌喷个水！", res[0].Username)
	// 调用发送函数
	Send2group(groupId, msg)
}

// 历史龙王
func CounterRankHistory(groupId int64) {
	// 查询数据
	res, err := models.MemberFindMany(bson.M{"group_id": groupId, "total_msg_count": bson.M{"$gt": 0}}, &options.FindOptions{Sort: bson.M{"total_msg_count": -1}})
	if err != nil {
		log.Println("CounterRankHistory(1):", err)
	}
	// 生成回复
	var msg string
	for i := range res {
		msg = msg + fmt.Sprintf("%d.%s:%d条\n", i+1, res[i].Username, res[i].TotalMsgCount)
	}
	msg = "历史骚话排行榜：\n" + msg + fmt.Sprintf("%s，永远滴龙王！", res[0].Username)
	// 调用发送函数
	Send2group(groupId, msg)
}
