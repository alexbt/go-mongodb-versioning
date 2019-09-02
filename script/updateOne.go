package script

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type updateOne struct {
	*meta
	Filter         bsonx.Doc
	Value          bsonx.Doc
	validCheckSums []string
}

func (op scriptWithOperation) WithUpdateOne(filter bsonx.Doc, value bsonx.Doc) Script {
	return updateOne{
		op.meta,
		filter,
		value,
		op.validCheckSums,
	}
}

func (u updateOne) execute(ctx context.Context, m *mongo.Database) {
	m.Collection(u.meta.collectionName).UpdateOne(ctx, u.Filter, u.Value)
}

func (u updateOne) getMeta() meta {
	return *u.meta
}

func (u updateOne) getContent() []interface{} {
	return []interface{}{u.Filter, u.Value}
}

func (u updateOne) getValidChecksums() []string {
	return u.validCheckSums
}
