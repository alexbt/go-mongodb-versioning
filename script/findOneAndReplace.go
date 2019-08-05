package script

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type findOneAndUpdate struct {
	*meta
	Filter bsonx.Doc
	Value  bsonx.Doc
}

func (op scriptImpl) FindOneAndUpdate(filter bsonx.Doc, value bsonx.Doc) Script {
	return findOneAndUpdate{op.meta, filter, value}
}

func (u findOneAndUpdate) execute(ctx context.Context, m *mongo.Database) {
	m.Collection(u.meta.collectionName).FindOneAndReplace(ctx, u.Filter, u.Value)
}

func (u findOneAndUpdate) getMeta() meta {
	return *u.meta
}
