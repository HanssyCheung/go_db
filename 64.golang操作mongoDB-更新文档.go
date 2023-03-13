package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//创建一个结构体

type Student3 struct {
	Name string
	Age  int
}

var client3 *mongo.Client

func initMDB3() {
	//设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//连接到MongoDB
	var err error
	client3, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//检查连接
	err = client3.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	println("Connected to MongoDB!")
}

func update() {
	initMDB3()
	ctx := context.TODO()
	defer client3.Disconnect(ctx)
	c := client3.Database("go_db").Collection("student")
	updateContent := bson.D{{"$set", bson.D{{"Name", "big tom"}, {"Age", 22}}}}
	ur, err := c.UpdateMany(ctx, bson.D{{"name", "tom"}}, updateContent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ur.ModifiedCount: %v\n", ur.ModifiedCount)
}

func main() {
	update()
}
