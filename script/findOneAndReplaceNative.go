package script

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type findOneAndUpdateNative struct {
	*meta
	Filter         string
	Value          string
	validCheckSums []string
}

func (op scriptWithOperation) WithFindOneAndUpdateNative(filter string, value string) Script {
	return findOneAndUpdateNative{op.meta, filter, value, op.validCheckSums}
}

func (u findOneAndUpdateNative) execute(ctx context.Context, m *mongo.Database) {
	var filter interface{}
	bson.UnmarshalJSON([]byte(u.Filter), &filter)

	var replace interface{}
	bson.UnmarshalJSON([]byte(u.Value), &replace)

	m.Collection(u.meta.collectionName).FindOneAndReplace(ctx, filter, replace)
}

func (u findOneAndUpdateNative) getMeta() meta {
	return *u.meta
}

func (u findOneAndUpdateNative) getValidChecksums() []string {
	return u.validCheckSums
}

func (u findOneAndUpdateNative) getContent() []interface{} {
	return []interface{}{u.Filter, u.Value}
}
