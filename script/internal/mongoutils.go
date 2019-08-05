package internal

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
