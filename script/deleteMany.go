package script

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type deleteMany struct {
	*meta
	Value          bsonx.Doc
	validCheckSums []string
}

func (op scriptWithOperation) WithWithDeleteMany(value bsonx.Doc) Script {
	return deleteMany{op.meta, value, op.validCheckSums}
}

func (u deleteMany) execute(ctx context.Context, m *mongo.Database) {
	m.Collection(u.meta.collectionName).DeleteMany(ctx, u.Value)
}

func (u deleteMany) getMeta() meta {
	return *u.meta
}

func (u deleteMany) getValidChecksums() []string {
	return u.validCheckSums
}

func (u deleteMany) getContent() []interface{} {
	return []interface{}{u.Value}
}
