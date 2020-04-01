package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type GroupMsg struct {
	GroupId   int64  `json:"group_id" bson:"group_id"`
	GroupName string `json:"group_name"`
	//LastModifiedAt int64 `json:"last_modified_at" json:"last_modified_at"`
	LastReplyAt int64 `json:"last_reply_at" bson:"last_reply_at"`
}

func GroupInsertOne(msg GroupMsg) (*mongo.InsertOneResult, error) {
	res, err := GroupColl.InsertOne(context.Background(), bson.M{
		"group_id":      msg.GroupId,
		"group_name":    msg.GroupName,
		"last_reply_at": time.Now().Unix(),
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

func GroupUpdateOne(filter, update bson.M) (*mongo.UpdateResult, error) {
	res, err := GroupColl.UpdateOne(context.Background(), filter, update)
	return res, err
}
func GroupUpdateMany(filter, update bson.M) (*mongo.UpdateResult, error) {
	res, err := GroupColl.UpdateMany(context.Background(), filter, update)
	return res, err
}
