package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Student4 struct {
	Name string
	Age  int
}

var client4 *mongo.Client

func initMDB4() {
	//设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//连接到MongoDB
	var err error
	client4, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//检查连接
	err = client4.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	println("Connected to MongoDB!")
}

//删除操作

func del() {
	initMDB4()
	c := client4.Database("go_db").Collection("student")
	ctx := context.TODO()

	dr, err := c.DeleteMany(ctx, bson.D{{"Name", "big tom"}})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ur.ModifiedCount: %v\n", dr.DeletedCount)
}

func main() {
	del()
}