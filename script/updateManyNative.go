package script

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type updateManyNative struct {
	*meta
	Filter         string
	Value          string
	validCheckSums []string
}

func (op scriptWithOperation) WithUpdateManyNative(filter string, value string) Script {
	return updateManyNative{
		op.meta,
		filter,
		value,
		op.validCheckSums,
	}
}

func (u updateManyNative) execute(ctx context.Context, m *mongo.Database) {
	var filter interface{}
	bson.UnmarshalJSON([]byte(u.Filter), &filter)

	var update interface{}
	bson.UnmarshalJSON([]byte(u.Value), &update)

	m.Collection(u.meta.collectionName).UpdateMany(ctx, filter, update)
}

func (u updateManyNative) getMeta() meta {
	return *u.meta
}

func (u updateManyNative) getContent() []interface{} {
	return []interface{}{u.Filter, u.Value}
}

func (u updateManyNative) getValidChecksums() []string {
	return u.validCheckSums
}
