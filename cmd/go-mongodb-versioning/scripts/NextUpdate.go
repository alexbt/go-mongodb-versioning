package scripts

import (
	"github.com/alexbt/go-mongodb-versioning/script"
)

func NextUpdate() script.Script {
	return script.NewScript().
		Meta("alexbt",
			"name2",
			"blah",
			"f2fa42fc5ba48b7bcbf6572c7dceb401").
		UpdateManyNative(
			"{}",
			`{$set: {"yo": "blah3"}}`)
}
