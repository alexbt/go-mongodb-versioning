package scripts

import (
	"github.com/alexbt/go-mongodb-versioning/script"
)

func NextNextUpdate() script.Script {
	return script.NewScript().
		Meta("alexbt", "nextNextUpdate", "blah", "4202f70124d2953fced73dbf3dd311e3").
		UpdateManyNative(
			`{"yo": "wrong"}`,
			`{$set: {"yo": "blah3"}}`)
}
