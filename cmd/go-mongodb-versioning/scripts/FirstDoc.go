package scripts

import (
	"github.com/alexbt/go-mongodb-versioning/script"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func FirstDoc() script.Script {
	return script.NewScript().
		WithMeta("alexbt", "firstdoc", "80d85b1494d0de47064015c86cd44eb0").
		WithUpdateMany(
			bsonx.Doc{},
			bsonx.Doc{
				{"$set", bsonx.Document(bsonx.Doc{
					{"yo", bsonx.String("blah55")},
				}),
				},
			})
}
