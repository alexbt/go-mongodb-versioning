package script

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type updateMany struct {
	*meta
	Filter bsonx.Doc
	Value  bsonx.Doc
}

func (op scriptImpl) UpdateMany(filter bsonx.Doc, value bsonx.Doc) Script {
	return updateMany{op.meta, filter, value}
}

func (u updateMany) execute(ctx context.Context, m *mongo.Database) {
	m.Collection(u.meta.collectionName).UpdateMany(ctx, u.Filter, u.Value)
}

func (u updateMany) getMeta() meta {
	return *u.meta
}
