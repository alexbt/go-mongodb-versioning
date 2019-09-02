package script

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type findOneAndReplace struct {
	*meta
	Filter         bsonx.Doc
	Value          bsonx.Doc
	validCheckSums []string
}

func (op scriptWithOperation) WithFindOneAndReplace(filter bsonx.Doc, value bsonx.Doc) Script {
	return findOneAndReplace{op.meta, filter, value, op.validCheckSums}
}

func (u findOneAndReplace) execute(ctx context.Context, m *mongo.Database) {
	m.Collection(u.meta.collectionName).FindOneAndUpdate(ctx, u.Filter, u.Value)
}

func (u findOneAndReplace) getMeta() meta {
	return *u.meta
}

func (u findOneAndReplace) getValidChecksums() []string {
	return u.validCheckSums
}

func (u findOneAndReplace) getContent() []interface{} {
	return []interface{}{u.Filter, u.Value}
}
