package internal

import (
	"context"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func OpenConnection(dbname string, url string) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	client.Connect(ctx)
	return client.Database(dbname)
}

func GetClient(url string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	return client
}

func BackupCollection(ctx context.Context, connection *mongo.Database, collName string) {
	var all interface{}
	bson.UnmarshalJSON([]byte(`{}`), &all)
	results, _ := connection.Collection(collName).Find(ctx, all)

	defer results.Close(ctx)

	var itemBson bson.M
	timstamps := time.Now().Format("20060102T150405")
	for results.Next(ctx) {
		results.Decode(&itemBson)
		connection.Collection(fmt.Sprintf("%s_%s", collName, timstamps)).InsertOne(ctx, itemBson)
	}

}
