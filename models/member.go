package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MemberMsg struct {
	GroupId   int64  `json:"group_id" bson:"group_id"`
	GroupName string `json:"group_name"`
	UserId    int64  `json:"user_id" bson:"user_id"`
	Username  string `json:"username" bson:"username"`

	LastMsgAt   int64 `json:"last_msg_at" bson:"last_msg_at"`
	LastReplyAt int64 `json:"last_reply_at" bson:"last_reply_at"`

	TodayMsgCount uint `json:"today_msg_count" bson:"today_msg_count"`
	TotalMsgCount uint `json:"total_msg_count" bson:"total_msg_count"`
}

func MemberInsertOne(msg MemberMsg) (*mongo.InsertOneResult, error) {
	res, err := MemberColl.InsertOne(context.Background(), bson.M{
		"group_id":   msg.GroupId,
		"group_name": msg.GroupName,
		"user_id":    msg.UserId,
		"username":   msg.Username,

		"last_msg_at":   time.Now().Unix(),
		"last_reply_at": time.Now().Unix(),

		"today_msg_count": 0,
		"total_msg_count": 0,
	})
	return res, err
}

func MemberFindOne(filter bson.M) (MemberMsg, error) {
	var msg MemberMsg
	err := MemberColl.FindOne(context.Background(), filter).Decode(&msg)
	return msg, err
}

func MemberFindOneAndUpdate(filter, update bson.M) (MemberMsg, error) {
	var msg MemberMsg
	err := MemberColl.FindOneAndUpdate(context.Background(), filter, update).Decode(&msg)
	return msg, err
}

func MemberFindMany(filter bson.M, option *options.FindOptions) ([]MemberMsg, error) {
	ctx := context.Background()
	cursor, err := MemberColl.Find(ctx, filter, option)
	if err != nil {
		return nil, err
	}
	// iterate through all documents
	var res []MemberMsg
	for cursor.Next(ctx) {
		var p MemberMsg
		// decode the document into given type
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		res = append(res, p)
	}
	return res, nil
}

func MemberUpdateOne(filter, update bson.M) (*mongo.UpdateResult, error) {
	res, err := MemberColl.UpdateOne(context.Background(), filter, update)
	return res, err
}
func MemberUpdateMany(filter, update bson.M) (*mongo.UpdateResult, error) {
	res, err := MemberColl.UpdateMany(context.Background(), filter, update)
	return res, err
}
