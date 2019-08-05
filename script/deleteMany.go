package script

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type deleteMany struct {
	*meta
	Value bsonx.Doc
}

func (op scriptImpl) DeleteMany(value bsonx.Doc) Script {
	return deleteMany{op.meta, value}
}

func (u deleteMany) execute(ctx context.Context, m *mongo.Database) {
	m.Collection(u.meta.collectionName).DeleteMany(ctx, u.Value)
}

func (u deleteMany) getMeta() meta {
	return *u.meta
}
