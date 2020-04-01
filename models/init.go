package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"qq_bot/config"
	"time"
)

var GroupColl *mongo.Collection
var MemberColl *mongo.Collection

func init() {
	// 初始化连接
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DatabaseURL))
	if err != nil {
		log.Fatal("Database init failed:", err)
	}
	// 检查数据库连接
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Database init failed:", err)
	}

	// 建立表
	GroupColl = client.Database(config.DatabaseName).Collection("group")
	MemberColl = client.Database(config.DatabaseName).Collection("member")
}
