package scripts

import (
	"github.com/alexbt/go-mongodb-versioning/script"
)

func MyFirst() script.Script {
	return script.NewScript().
		Meta("alexbt", "name", "blah", "0226183a64c83de70570cd422e11e9cf").
		UpdateManyNative("{}",
			`{$set: {"yo": "blah2"}}`)
}
