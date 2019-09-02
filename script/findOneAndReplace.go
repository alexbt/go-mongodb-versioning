package script

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type findOneAndUpdate struct {
	*meta
	Filter         bsonx.Doc
	Value          bsonx.Doc
	validCheckSums []string
}

func (op scriptWithOperation) WithFindOneAndUpdate(filter bsonx.Doc, value bsonx.Doc) Script {
	return findOneAndUpdate{op.meta, filter, value, op.validCheckSums}
}

func (u findOneAndUpdate) execute(ctx context.Context, m *mongo.Database) {
	m.Collection(u.meta.collectionName).FindOneAndReplace(ctx, u.Filter, u.Value)
}

func (u findOneAndUpdate) getMeta() meta {
	return *u.meta
}

func (u findOneAndUpdate) getValidChecksums() []string {
	return u.validCheckSums
}

func (u findOneAndUpdate) getContent() []interface{} {
	return []interface{}{u.Filter, u.Value}
}
