package main

import (
	"github.com/alexbt/go-mongodb-versioning/cmd/go-mongodb-versioning/scripts"
	"github.com/alexbt/go-mongodb-versioning/script"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	runner := script.NewRunner(
		"go-mongo-versioning",
		"dummy")

	changeSets := script.NewChangeSets(
		scripts.MyFirst(),
		scripts.NextUpdate(),
		scripts.NextNextUpdate(),
		scripts.FirstInsert(),
		scripts.SecondInsert(),
		scripts.FirstDoc(),
	)

	runner.Execute(changeSets)
}
