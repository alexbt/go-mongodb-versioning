package script

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type replaceOneNative struct {
	*meta
	Filter         string
	Value          string
	validCheckSums []string
}

func (op scriptWithOperation) WithReplaceOneNative(filter string, value string) Script {
	return replaceOneNative{
		op.meta,
		filter,
		value,
		op.validCheckSums,
	}
}

func (u replaceOneNative) execute(ctx context.Context, m *mongo.Database) {
	var filter interface{}
	bson.UnmarshalJSON([]byte(u.Filter), &filter)

	var update interface{}
	bson.UnmarshalJSON([]byte(u.Value), &update)

	m.Collection(u.meta.collectionName).ReplaceOne(ctx, filter, update)
}

func (u replaceOneNative) getMeta() meta {
	return *u.meta
}

func (u replaceOneNative) getContent() []interface{} {
	return []interface{}{u.Filter, u.Value}
}

func (u replaceOneNative) getValidChecksums() []string {
	return u.validCheckSums
}
