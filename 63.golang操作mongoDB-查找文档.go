package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client2 *mongo.Client

func initMDB2() {
	//设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//连接到MongoDB
	var err error
	client2, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//检查连接
	err = client2.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	println("Connected to MongoDB!")
}

func find() {
	initMDB2()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := client2.Database("go_db").Collection("student")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("result:%v\n", result)
		fmt.Printf("result.Map():%v\n", result.Map()["name"])
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	find()
}
