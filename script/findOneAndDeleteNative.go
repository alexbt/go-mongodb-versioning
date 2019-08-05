package script

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type findOneAndDeleteNative struct {
	*meta
	Filter string
}

func (op scriptImpl) FindOneAndDeleteNative(filter string) Script {
	return findOneAndDeleteNative{op.meta, filter}
}

func (u findOneAndDeleteNative) execute(ctx context.Context, m *mongo.Database) {
	var filter interface{}
	bson.UnmarshalJSON([]byte(u.Filter), &filter)

	m.Collection(u.meta.collectionName).FindOneAndDelete(ctx, filter)
}

func (u findOneAndDeleteNative) getMeta() meta {
	return *u.meta
}
