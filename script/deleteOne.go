package script

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type deleteOne struct {
	*meta
	Value bsonx.Doc
}

func (op scriptImpl) DeleteOne(value bsonx.Doc) Script {
	return deleteOne{op.meta, value}
}

func (u deleteOne) execute(ctx context.Context, m *mongo.Database) {
	m.Collection(u.meta.collectionName).DeleteOne(ctx, u.Value)
}

func (u deleteOne) getMeta() meta {
	return *u.meta
}
