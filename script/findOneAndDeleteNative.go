package script

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type findOneAndDeleteNative struct {
	*meta
	Filter         string
	validCheckSums []string
}

func (op scriptWithOperation) WithFindOneAndDeleteNative(filter string) Script {
	return findOneAndDeleteNative{op.meta, filter, op.validCheckSums}
}

func (u findOneAndDeleteNative) execute(ctx context.Context, m *mongo.Database) {
	var filter interface{}
	bson.UnmarshalJSON([]byte(u.Filter), &filter)

	m.Collection(u.meta.collectionName).FindOneAndDelete(ctx, filter)
}

func (u findOneAndDeleteNative) getMeta() meta {
	return *u.meta
}

func (u findOneAndDeleteNative) getValidChecksums() []string {
	return u.validCheckSums
}

func (u findOneAndDeleteNative) getContent() []interface{} {
	return []interface{}{u.Filter}
}
