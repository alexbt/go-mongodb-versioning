package scripts

import (
	"github.com/alexbt/go-mongodb-versioning/script"
)

func SecondInsert() script.Script {
	return script.NewScript().
		WithMeta(
			"alexbt",
			"insert2",
			"blah", "bbb47067604603eddfa514fbcb30533b").
		WithInsertOneNative(
			`{"coucou": "damn"}`)
}
