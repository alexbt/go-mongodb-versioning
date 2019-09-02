package script

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type updateOneNative struct {
	*meta
	Filter         string
	Value          string
	validCheckSums []string
}

func (op scriptWithOperation) WithUpdateOneNative(filter string, value string) Script {
	return updateOneNative{
		op.meta,
		filter,
		value,
		op.validCheckSums}
}

func (u updateOneNative) execute(ctx context.Context, m *mongo.Database) {
	var filter interface{}
	bson.UnmarshalJSON([]byte(u.Filter), &filter)

	var update interface{}
	bson.UnmarshalJSON([]byte(u.Value), &update)

	m.Collection(u.meta.collectionName).UpdateOne(ctx, filter, update)
}

func (u updateOneNative) getMeta() meta {
	return *u.meta
}

func (u updateOneNative) getContent() []interface{} {
	return []interface{}{u.Filter, u.Value}
}

func (u updateOneNative) getValidChecksums() []string {
	return u.validCheckSums
}
