package scripts

import (
	"github.com/alexbt/go-mongodb-versioning/script"
)

func FirstInsert() script.Script {
	return script.NewScript().
		WithMeta("alexbt", "insert", "blah", "88783e207575d858cc48450468500970").
		WithUpdateManyNative(
			"{}",
			`{$set: {"yo": "blah3"}}`)
}
