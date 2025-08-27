package main

import (
	"github.com/coalaura/plain"
)

func SelectDatabase(databases []*Database) *Database {
	index, err := log.Select("Database > ", plain.AsStrings(databases))
	if err != nil {
		return nil
	}

	return databases[index]
}

func SelectType() string {
	options := []string{"user", "character"}

	index, err := log.Select("Type > ", options)
	if err != nil {
		return ""
	}

	return options[index]
}
