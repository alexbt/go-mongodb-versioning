package script

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type deleteManyNative struct {
	*meta
	Value string
}

func (op scriptImpl) DeleteManyNative(value string) Script {
	return deleteManyNative{op.meta, value}
}

func (u deleteManyNative) execute(ctx context.Context, m *mongo.Database) {
	var remove interface{}
	bson.UnmarshalJSON([]byte(u.Value), &remove)
	m.Collection(u.meta.collectionName).DeleteMany(ctx, remove)
}

func (u deleteManyNative) getMeta() meta {
	return *u.meta
}
