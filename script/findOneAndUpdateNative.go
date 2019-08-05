package script

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type findOneAndReplaceNative struct {
	*meta
	Filter string
	Value  string
}

func (op scriptImpl) FindOneAndReplaceNative(filter string, value string) Script {
	return findOneAndReplaceNative{op.meta, filter, value}
}

func (u findOneAndReplaceNative) execute(ctx context.Context, m *mongo.Database) {
	var filter interface{}
	bson.UnmarshalJSON([]byte(u.Filter), &filter)

	var update interface{}
	bson.UnmarshalJSON([]byte(u.Value), &update)

	m.Collection(u.meta.collectionName).FindOneAndUpdate(ctx, filter, update)
}

func (u findOneAndReplaceNative) getMeta() meta {
	return *u.meta
}
