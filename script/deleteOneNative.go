package script

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type deleteOneNative struct {
	*meta
	Value          string
	validCheckSums []string
}

func (op scriptWithOperation) WithDeleteOneNative(value string) Script {
	return deleteOneNative{op.meta, value, op.validCheckSums}
}

func (u deleteOneNative) execute(ctx context.Context, m *mongo.Database) {
	var remove interface{}
	bson.UnmarshalJSON([]byte(u.Value), &remove)
	m.Collection(u.meta.collectionName).DeleteOne(ctx, remove)
}

func (u deleteOneNative) getMeta() meta {
	return *u.meta
}

func (u deleteOneNative) getValidChecksums() []string {
	return u.validCheckSums
}

func (u deleteOneNative) getContent() []interface{} {
	return []interface{}{u.Value}
}
