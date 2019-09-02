package script

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type findOneAndDelete struct {
	*meta
	Filter         bsonx.Doc
	validCheckSums []string
}

func (op scriptWithOperation) WithFindOneAndDelete(filter bsonx.Doc) Script {
	return findOneAndDelete{op.meta, filter, op.validCheckSums}
}

func (u findOneAndDelete) execute(ctx context.Context, m *mongo.Database) {
	m.Collection(u.meta.collectionName).FindOneAndDelete(ctx, u.Filter)
}

func (u findOneAndDelete) getMeta() meta {
	return *u.meta
}

func (u findOneAndDelete) getValidChecksums() []string {
	return u.validCheckSums
}

func (u findOneAndDelete) getContent() []interface{} {
	return []interface{}{u.Filter}
}
