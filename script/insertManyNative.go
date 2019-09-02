package script

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type insertManyNative struct {
	*meta
	Value          string
	validCheckSums []string
}

func (op scriptWithOperation) WithInsertManyNative(value string) Script {
	return insertManyNative{op.meta, value, op.validCheckSums}
}

func (u insertManyNative) execute(ctx context.Context, m *mongo.Database) {
	var insert []interface{}
	bson.UnmarshalJSON([]byte(u.Value), &insert)
	m.Collection(u.meta.collectionName).InsertMany(ctx, insert)
}

func (u insertManyNative) getMeta() meta {
	return *u.meta
}

func (u insertManyNative) getValidChecksums() []string {
	return u.validCheckSums
}

func (u insertManyNative) getContent() []interface{} {
	return []interface{}{u.Value}
}
