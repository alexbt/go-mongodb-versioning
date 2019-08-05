package script

import (
	"context"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"go.mongodb.org/mongo-driver/mongo"
)

type insertOne struct {
	*meta
	Value bsonx.Doc
}

func (op scriptImpl) InsertOne(value bsonx.Doc) Script {
	return insertOne{op.meta, value}
}

func (u insertOne) execute(ctx context.Context, m *mongo.Database) {
	m.Collection(u.meta.collectionName).InsertOne(ctx, u.Value)
}

func (u insertOne) getMeta() meta {
	return *u.meta
}
