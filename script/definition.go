package script

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type Script interface {
	execute(ctx context.Context, m *mongo.Database)
	getMeta() meta
}
type changeSets []Script

type scriptWithOperation struct {
	*meta
}

type scriptImplWithMeta struct {
}

type meta struct {
	uniqueName     string
	author         string
	collectionName string
	validCheckSums []string
}

func NewChangeSets(scripts ...Script) changeSets {
	if scripts == nil {
		panic(errors.New("Cannot initialize with nil changeSet"))
	}
	return scripts
}

func NewScript() *scriptImplWithMeta {
	return &scriptImplWithMeta{}
}

func (op *scriptImplWithMeta) WithMeta(author string, uniqueName string, collectionName string, validCheckSums ...string) *scriptWithOperation {
	return &scriptWithOperation{
		&meta{
			uniqueName:     uniqueName,
			author:         author,
			collectionName: collectionName,
			validCheckSums: validCheckSums,
		},
	}
}
