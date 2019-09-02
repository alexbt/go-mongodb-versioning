package script

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type replaceOne struct {
	*meta
	Filter         bsonx.Doc
	Value          bsonx.Doc
	validCheckSums []string
}

func (op scriptWithOperation) WithReplaceOne(filter bsonx.Doc, value bsonx.Doc) Script {
	return replaceOne{op.meta, filter, value, op.validCheckSums}
}

func (u replaceOne) execute(ctx context.Context, m *mongo.Database) {
	m.Collection(u.meta.collectionName).ReplaceOne(ctx, u.Filter, u.Value)
}

func (u replaceOne) getMeta() meta {
	return *u.meta
}

func (u replaceOne) getContent() []interface{} {
	return []interface{}{u.Filter, u.Value}
}

func (u replaceOne) getValidChecksums() []string {
	return u.validCheckSums
}
