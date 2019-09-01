package script

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/alexbt/go-mongodb-versioning/script/internal"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
)

type runner struct {
	dbName string
	client *mongo.Client
}

func NewRunner(DbName string, connection string) *runner {
	return &runner{
		dbName: DbName,
		client: internal.GetClient(connection),
	}
}

func NewRunnerWithClient(DbName string, client *mongo.Client) *runner {
	return &runner{
		dbName: DbName,
		client: client,
	}
}

func (r runner) Execute(changeSets changeSets) {
	mymap := make(map[string]bool)
	for _, v := range changeSets {
		if mymap[v.getMeta().uniqueName] {
			panic(errors.New("Duplicate unique name"))
		}
		mymap[v.getMeta().uniqueName] = true
	}

	ctx := context.Background()

	for _, v := range changeSets {
		bytesForMd5, _ := bson.Marshal(v)
		meta := internal.ScriptMeta{
			Operation:      reflect.TypeOf(v).Name(),
			UniqueName:     v.getMeta().uniqueName,
			Author:         v.getMeta().author,
			Md5:            fmt.Sprintf("%x", md5.Sum(bytesForMd5)),
			ValidCheckSums: v.getMeta().validCheckSums,
		}

		db := r.client.Database(r.dbName)
		if meta.HasRun(db, ctx) {
			fmt.Println(fmt.Sprintf("Skipping %s-%s.%s (md5:%s - valids:%s)", meta.Author, meta.UniqueName, meta.Operation, meta.Md5, meta.ValidCheckSums))
			continue
		}

		fmt.Printf("Executing %s-%s.%s (md5:%s - valids:%s)", meta.Author, meta.UniqueName, meta.Operation, meta.Md5, meta.ValidCheckSums)

		internal.BackupCollection(ctx, db, v.getMeta().collectionName)
		fmt.Printf(" ...collection backed up")

		v.execute(ctx, db)
		fmt.Printf(" ...Executed")

		meta.Save(db, ctx)
		fmt.Printf(" ...Recorded")
		fmt.Println()
	}
}
