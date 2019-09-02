package script

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type insertMany struct {
	*meta
	Value          []bsonx.Doc
	validCheckSums []string
}

func (op scriptWithOperation) WithInsertMany(value []bsonx.Doc) Script {
	return insertMany{op.meta, value, op.validCheckSums}
}

func (u insertMany) execute(ctx context.Context, m *mongo.Database) {
	values := make([]interface{}, len(u.Value))

	for i, v := range u.Value {
		values[i] = v
	}

	m.Collection(u.meta.collectionName).InsertMany(ctx, values)
}

func (u insertMany) getMeta() meta {
	return *u.meta
}

func (u insertMany) getValidChecksums() []string {
	return u.validCheckSums
}

func (u insertMany) getContent() []interface{} {
	return []interface{}{u.Value}
}
