package script

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/alexbt/go-mongodb-versioning/script/internal"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

type runner struct {
	DbName     string
	Connection string
}

func NewRunner(DbName string, connection string) *runner {
	return &runner{
		DbName:     DbName,
		Connection: connection,
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

	connection := internal.OpenConnection(r.DbName, r.Connection)
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
		if meta.HasRun(connection, ctx) {
			fmt.Println(fmt.Sprintf("Skipping %s-%s.%s (md5:%s - valids:%s)", meta.Author, meta.UniqueName, meta.Operation, meta.Md5, meta.ValidCheckSums))
			continue
		}

		fmt.Printf("Executing %s-%s.%s (md5:%s - valids:%s)", meta.Author, meta.UniqueName, meta.Operation, meta.Md5, meta.ValidCheckSums)
		v.execute(ctx, connection)
		fmt.Printf(" ...Executed")

		meta.Save(connection, ctx)
		fmt.Printf(" ...Recorded")
	}
}
