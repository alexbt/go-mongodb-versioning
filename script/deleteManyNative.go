package script

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type deleteManyNative struct {
	*meta
	Value          string
	validCheckSums []string
}

func (op scriptWithOperation) WithDeleteManyNative(value string) Script {
	return deleteManyNative{op.meta, value, op.validCheckSums}
}

func (u deleteManyNative) execute(ctx context.Context, m *mongo.Database) {
	var remove interface{}
	bson.UnmarshalJSON([]byte(u.Value), &remove)
	m.Collection(u.meta.collectionName).DeleteMany(ctx, remove)
}

func (u deleteManyNative) getMeta() meta {
	return *u.meta
}

func (u deleteManyNative) getValidChecksums() []string {
	return u.validCheckSums
}

func (u deleteManyNative) getContent() []interface{} {
	return []interface{}{u.Value}
}
