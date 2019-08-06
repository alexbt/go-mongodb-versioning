package scripts

import (
	"github.com/alexbt/go-mongodb-versioning/script"
)

func MyReplace() script.Script {
	return script.NewScript().
		WithMeta("alexbt", "insert2", "blah", "bbb47067604603eddfa514fbcb30533b").
		WithReplaceOneNative(
			`{}`,
			`{"coucou": "damn"}`)
}
