package main

import (
	"os"

	"github.com/coalaura/plain"
)

var log = plain.New(os.Stdout)

func main() {
	log.Println("Loading config...")

	config, err := LoadConfig()
	log.MustFail(err)

	log.Println("Loading databases...")

	databases, err := LoadDatabases(config)
	log.MustFail(err)

	database := SelectDatabase(databases)
	if database == nil {
		return
	}

	log.Printf("Connecting to %d...\n", database.ID)

	err = database.Connect()
	log.MustFail(err)

	loop(database)
}

func loop(database *Database) {
	for {
		var (
			typ   = SelectType()
			where string
			err   error
		)

		for where == "" {
			where, _ = log.Read(os.Stdin, "Where > ", 64)
		}

		log.Println("Retrieving...")

		switch typ {
		case "user":
			err = HandleUser(database, where)
		case "character":
			err = HandleCharacter(database, where)
		}

		if err != nil {
			log.Warnln(err.Error())
		}
	}
}
