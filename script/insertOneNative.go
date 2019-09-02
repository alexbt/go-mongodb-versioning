package script

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type insertOneNative struct {
	*meta
	Value          string
	validCheckSums []string
}

func (op scriptWithOperation) WithInsertOneNative(value string) Script {
	return insertOneNative{op.meta, value, op.validCheckSums}
}

func (u insertOneNative) execute(ctx context.Context, m *mongo.Database) {
	var insert interface{}
	bson.UnmarshalJSON([]byte(u.Value), &insert)

	m.Collection(u.meta.collectionName).InsertOne(ctx, insert)
}

func (u insertOneNative) getMeta() meta {
	return *u.meta
}

func (u insertOneNative) getValidChecksums() []string {
	return u.validCheckSums
}

func (u insertOneNative) getContent() []interface{} {
	return []interface{}{u.Value}
}
