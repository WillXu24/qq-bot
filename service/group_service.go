package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"qq_bot/models"
	"qq_bot/utils"
	"qq_bot/views"
)

// MsgCounter 用于更新消息计数
func MsgCounter(msg views.PostMsg) error {
	// 查找记录
	filter := bson.M{"group_id": msg.GroupId, "user_id": msg.UserId}
	res, err := models.GroupFindOne(filter)
	if err != nil {
		// 不存在则创建,返回
		_, err := models.GroupInsertOne(models.GroupMsg{
			GroupId:     msg.GroupId,
			UserId:      msg.UserId,
			Username:    msg.Sender.Nickname,
			LastMsgAt:   msg.Time,
			LastReplyAt: 0,
		})
		return err
	}
	// 检查日期
	timeZero := utils.GetTimeZero()
	if res.LastMsgAt < timeZero {
		// 新的一天则相应更新，返回
		_, err := models.GroupUpdateOne(filter, bson.M{
			"$set": bson.M{
				"last_msg_at":     msg.Time,
				"today_msg_count": 1,
			},
			"$inc": bson.M{
				"total_msg_count": 1,
			}})
		return err
	}
	// 同一天则默认更新，返回
	_, err = models.GroupUpdateOne(filter, bson.M{
		"$set": bson.M{
			"last_msg_at": msg.Time,
		},
		"$inc": bson.M{
			"today_msg_count": 1,
			"total_msg_count": 1,
		}})
	return err
}

// 获取今日龙王：
// 今日骚话排行榜：\n
// 1.某某某:多少条\n
// 2.某某某：多少条\n
//【某某某】快给大家喷个水！
func GetDragonToday(groupId int64) error {
	// 查询数据
	res, err := models.GroupFindMany(bson.M{"group_id": groupId}, &options.FindOptions{Sort: bson.M{"today_msg_count": -1}})
	if err != nil {
		return err
	}
	// 生成回复
	var msg string
	for i := range res {
		msg = msg + fmt.Sprintf("%d.%s:%d条\n", i+1, res[i].Username, res[i].TodayMsgCount)
	}
	msg = "今日骚话排行榜：\n" + msg + fmt.Sprintf("%s，快给大家喷个水！", res[0].Username)
	// 发送回复
	return Send2group(groupId, msg)
}

// 获取历史龙王
// 历史骚话排行榜：\n
// 1.某某某:多少条\n
// 2.某某某：多少条\n
// 某某某，永远滴龙王！
func GetDragonHistory(groupId int64) error {
	// 查询数据
	res, err := models.GroupFindMany(bson.M{"group_id": groupId}, &options.FindOptions{Sort: bson.M{"today_msg_count": -1}})
	if err != nil {
		return err
	}
	// 生成回复
	var msg string
	for i := range res {
		msg = msg + fmt.Sprintf("%d.%s:%d条\n", i+1, res[i].Username, res[i].TodayMsgCount)
	}
	msg = "历史骚话排行榜：\n" + msg + fmt.Sprintf("%s，永远滴龙王！", res[0].Username)
	// 发送回复
	return Send2group(groupId, msg)
}
