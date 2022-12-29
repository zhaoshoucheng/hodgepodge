package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	mongoConn := "mongodb://127.0.0.1:27017"
	mongoDbName := "test"

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoConn))
	if err != nil {
		fmt.Println("NewClient err", err)
		return
	}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Connect", err)
		return
	}

	data := []interface{}{
		bson.M{
			"m_type":             "round_robin",
			"m_lastModifiedDate": "2017-02-09 09:10:10.129302",
			"m_creationDate":     "2017-02-08 20:54:35.716960",
			"m_target":           "",
			"m_name":             "round",
		},
		bson.M{
			"m_type":             "ip_hash",
			"m_lastModifiedDate": "2017-02-09 09:10:10.129302",
			"m_creationDate":     "2017-02-08 20:54:35.716960",
			"m_target":           "",
			"m_name":             "ip_hash",
		},
		bson.M{
			"m_type":             "least_conn",
			"m_lastModifiedDate": "2017-02-09 09:10:10.129302",
			"m_creationDate":     "2017-02-08 20:54:35.716960",
			"m_target":           "",
			"m_name":             "leastconn",
		},
		bson.M{
			"m_type":             "sticky",
			"m_lastModifiedDate": "2017-02-09 09:10:10.129302",
			"m_creationDate":     "2017-02-08 20:54:35.716960",
			"m_target":           "",
			"m_name":             "sticky",
		},
		bson.M{
			"m_type":             "round_robin",
			"m_lastModifiedDate": "2017-02-09 09:10:10.129302",
			"m_creationDate":     "2017-02-08 20:54:35.716960",
			"m_target":           "",
			"m_name":             "test_strageties",
		},
	}

	insert, err := client.Database(mongoDbName).Collection("strategies").InsertMany(
		context.Background(), data)
	if err != nil {
		fmt.Println("nginx_global_conf insert err", err)
		return
	}
	_ = insert
	count, err := client.Database(mongoDbName).Collection("strategies").CountDocuments(ctx, bson.M{})
	if err != nil {
		fmt.Println("CountDocuments", err)
		return
	}
	fmt.Println(count)

	fmt.Println("done!")
}
