package main

import (
	"fmt"
	"strings"

	"github.com/coalaura/logger/plain"
)

var log = plain.New()

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

	defer database.Close()

	for {
		license := prompt("License > ")

		if license == "" {
			continue
		} else if strings.EqualFold(string(license[0]), "q") {
			return
		} else if !strings.HasPrefix(license, "license:") {
			log.Warnln("Invalid license identifier")

			continue
		}

		user, err := database.Find(license)
		if err != nil {
			log.Warnln(err.Error())

			continue
		}

		log.Printf("\n%s\n", user.String())
	}
}

func prompt(msg string) string {
	var input string

	log.Print(msg)
	fmt.Scanln(&input)

	return strings.TrimSpace(input)
}
