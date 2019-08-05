package script

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type insertOneNative struct {
	*meta
	Value string
}

func (op scriptImpl) InsertOneNative(value string) Script {
	return insertOneNative{op.meta, value}
}

func (u insertOneNative) execute(ctx context.Context, m *mongo.Database) {
	var insert interface{}
	bson.UnmarshalJSON([]byte(u.Value), &insert)

	m.Collection(u.meta.collectionName).InsertOne(ctx, insert)
}

func (u insertOneNative) getMeta() meta {
	return *u.meta
}
