package script

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Script interface {
	execute(ctx context.Context, m *mongo.Database)
	getMeta() meta
}
type scriptImpl struct {
	*meta
}

type changeSets []Script

func NewChangeSets(scripts ...Script) changeSets {
	return []Script{}
}

type meta struct {
	uniqueName     string
	author         string
	collectionName string
	validCheckSums []string
}

//func NewMeta(author string, uniqueName string, collectionName string, validCheckSums ...string) *meta {
//	return &meta{
//		uniqueName:     uniqueName,
//		author:         author,
//		collectionName: collectionName,
//		validCheckSums: validCheckSums,
//	}
//}

func NewScript() *scriptImpl {
	return &scriptImpl{}
}

func (op *scriptImpl) Meta(author string, uniqueName string, collectionName string, validCheckSums ...string) *scriptImpl {
	op.meta = &meta{
		uniqueName:     uniqueName,
		author:         author,
		collectionName: collectionName,
		validCheckSums: validCheckSums,
	}
	return op
}
