package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//创建一个结构体

type Student1 struct {
	Name string
	Age  int
}

var client1 *mongo.Client

func initMDB1() {
	//设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//连接到MongoDB
	var err error
	client1, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//检查连接
	err = client1.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	println("Connected to MongoDB!")
}

//添加单个文档
//使用collection.InsertOne()方法插入一条文档记录:
func insertOne(s Student1) {
	initMDB1()
	collection := client1.Database("go_db").Collection("student")
	insertResult, err := collection.InsertOne(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
	}
	println("Inserted a single document", insertResult.InsertedID)
}

func main() {
	s := Student1{Name: "tom", Age: 20}
	insertOne(s)
}
