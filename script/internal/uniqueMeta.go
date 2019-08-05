package internal

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLL_NAME = "changeslog"

type ScriptMeta struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`
	Operation      string             `bson:"Operation"`
	UniqueName     string             `bson:"UniqueName"`
	Author         string             `bson:"Author"`
	Md5            string             `bson:"Md5"`
	ValidCheckSums []string           `bson:"validChecksums"`
}

type ScriptMetaSearchUnique struct {
	UniqueName string `bson:"UniqueName"`
	Author     string `bson:"Author"`
}

func (u ScriptMeta) HasRun(m *mongo.Database, ctx context.Context) bool {
	UniqueMetaSearch := ScriptMetaSearchUnique{
		UniqueName: u.UniqueName,
		Author:     u.Author,
	}
	one := m.Collection(COLL_NAME).FindOne(ctx, UniqueMetaSearch)
	if one.Err() == nil {
		return u.hasValidChecksum(one)
	}
	return true
}

func (u ScriptMeta) hasValidChecksum(result *mongo.SingleResult) bool {
	raws, _ := result.DecodeBytes()
	var fetched ScriptMeta
	result.Decode(&fetched)

	if raws != nil && fetched.Md5 == u.Md5 {
		return true
	} else if raws != nil {
		for _, v := range u.ValidCheckSums {
			if fetched.Md5 == v {
				return true
			}
		}

		er := fmt.Sprintf("Failing on %s-%s.%s (%s) - found %s", u.Author, u.UniqueName, u.Operation, u.Md5, fetched.Md5)
		fmt.Println(er)
		panic(errors.New(er))
	}

	return false
}

func (u ScriptMeta) Save(m *mongo.Database, ctx context.Context) (*mongo.InsertOneResult, error) {
	oneResult, er := m.Collection(COLL_NAME).InsertOne(ctx, u)
	return oneResult, er
}
