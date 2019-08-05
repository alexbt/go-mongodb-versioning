package script

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type deleteOneNative struct {
	*meta
	Value string
}

func (op scriptImpl) DeleteOneNative(value string) Script {
	return deleteOneNative{op.meta, value}
}

func (u deleteOneNative) execute(ctx context.Context, m *mongo.Database) {
	var remove interface{}
	bson.UnmarshalJSON([]byte(u.Value), &remove)
	m.Collection(u.meta.collectionName).DeleteOne(ctx, remove)
}

func (u deleteOneNative) getMeta() meta {
	return *u.meta
}
