package scripts

import (
	"github.com/alexbt/go-mongodb-versioning/script"
)

func FirstInsert() script.Script {
	return script.NewScript().
		Meta("alexbt", "insert", "blah", "88783e207575d858cc48450468500970").
		UpdateManyNative(
			"{}",
			`{$set: {"yo": "blah3"}}`)
}
