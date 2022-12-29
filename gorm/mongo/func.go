package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func CreateIndexDemo(conn mongo.Collection) error {
	indexView := conn.Indexes()
	models := []mongo.IndexModel{
		{
			Keys: bson.D{{"pkg_version", -1}},
		},
	}
	opts := options.CreateIndexes().SetMaxTime(2 * time.Second)
	names, err := indexView.CreateMany(context.TODO(), models, opts)
	if err != nil {
		return err
	}
	fmt.Println(names)
	return nil
}
