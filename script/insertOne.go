package script

import (
	"context"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"go.mongodb.org/mongo-driver/mongo"
)

type insertOne struct {
	*meta
	Value          bsonx.Doc
	validCheckSums []string
}

func (op scriptWithOperation) WithInsertOne(value bsonx.Doc) Script {
	return insertOne{op.meta, value, op.validCheckSums}
}

func (u insertOne) execute(ctx context.Context, m *mongo.Database) {
	m.Collection(u.meta.collectionName).InsertOne(ctx, u.Value)
}

func (u insertOne) getMeta() meta {
	return *u.meta
}

func (u insertOne) getValidChecksums() []string {
	return u.validCheckSums
}

func (u insertOne) getContent() []interface{} {
	return []interface{}{u.Value}
}
