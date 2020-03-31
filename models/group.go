package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GroupMsg struct {
	GroupId  int64  `json:"group_id" bson:"group_id"`
	UserId   int64  `json:"user_id" bson:"user_id"`
	Username string `json:"username" bson:"username"`

	LastMsgAt   int64 `json:"last_msg_at" bson:"last_msg_at"`
	LastReplyAt int64 `json:"last_reply_at" bson:"last_reply_at"`

	TodayMsgCount uint `json:"today_msg_count" bson:"today_msg_count"`
	TotalMsgCount uint `json:"total_msg_count" bson:"total_msg_count"`
}

func GroupInsertOne(msg GroupMsg) (*mongo.InsertOneResult, error) {
	res, err := GroupColl.InsertOne(context.Background(), bson.M{
		"group_id": msg.GroupId,
		"user_id":  msg.UserId,
		"username": msg.Username,

		"last_msg_at":   msg.LastMsgAt,
		"last_reply_at": msg.LastReplyAt,

		"today_msg_count": 1,
		"total_msg_count": 1,
	})
	return res, err
}

func GroupFindOne(filter bson.M) (GroupMsg, error) {
	var msg GroupMsg
	err := GroupColl.FindOne(context.Background(), filter).Decode(&msg)
	return msg, err
}

func GroupFindOneAndUpdate(filter, update bson.M) (GroupMsg, error) {
	var msg GroupMsg
	err := GroupColl.FindOneAndUpdate(context.Background(), filter, update).Decode(&msg)
	return msg, err
}

func GroupFindMany(filter bson.M, option *options.FindOptions) ([]GroupMsg, error) {
	ctx := context.Background()
	cursor, err := GroupColl.Find(ctx, filter, option)
	if err != nil {
		return nil, err
	}
	// iterate through all documents
	var res []GroupMsg
	for cursor.Next(ctx) {
		var p GroupMsg
		// decode the document into given type
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		res = append(res, p)
	}
	return res, nil
}

func GroupUpdateOne(filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	res, err := GroupColl.UpdateOne(context.Background(), filter, update)
	return res, err
}
