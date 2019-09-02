package script

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type deleteOne struct {
	*meta
	Value          bsonx.Doc
	validCheckSums []string
}

func (op scriptWithOperation) WithDeleteOne(value bsonx.Doc) Script {
	return deleteOne{op.meta, value, op.validCheckSums}
}

func (u deleteOne) execute(ctx context.Context, m *mongo.Database) {
	m.Collection(u.meta.collectionName).DeleteOne(ctx, u.Value)
}

func (u deleteOne) getMeta() meta {
	return *u.meta
}

func (u deleteOne) getValidChecksums() []string {
	return u.validCheckSums
}

func (u deleteOne) getContent() []interface{} {
	return []interface{}{u.Value}
}
