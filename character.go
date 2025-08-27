package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Character struct {
	ID             int64   `gorm:"column:character_id;<-:false"`
	FirstName      *string `gorm:"column:first_name;<-:false"`
	LastName       *string `gorm:"column:last_name;<-:false"`
	DateOfBirth    *string `gorm:"column:date_of_birth;<-:false"`
	Gender         *int64  `gorm:"column:gender;<-:false"`
	Cash           int64   `gorm:"column:cash;<-:false"`
	Bank           int64   `gorm:"column:bank;<-:false"`
	LastSeen       *int64  `gorm:"column:last_seen;<-:false"`
	Playtime       int64   `gorm:"column:playtime;<-:false"`
	JobName        *string `gorm:"column:job_name;<-:false"`
	DepartmentName *string `gorm:"column:department_name;<-:false"`
	PositionName   *string `gorm:"column:position_name;<-:false"`
}

func HandleCharacter(database *Database, where string) error {
	var (
		characters []Character
		err        error
	)

	if IsValidInteger(where) {
		characters, err = database.FindCharactersByID(where)
	} else if IsValidLicense(where) {
		characters, err = database.FindCharactersByLicense(where)
	} else {
		err = errors.New("invalid where (allowed: id, license)")
	}

	if err != nil {
		return err
	}

	for _, character := range characters {
		log.Println()
		log.Println(character)
	}

	return nil
}

func (c *Character) GetFullName() string {
	var name string

	if c.FirstName == nil && c.LastName == nil {
		return "n/a"
	}

	if c.FirstName != nil {
		name = *c.FirstName
	}

	if c.LastName != nil {
		if name != "" {
			name += " "
		}

		name += *c.LastName
	}

	return name
}

func (c *Character) GetGender() string {
	if c.Gender != nil {
		switch *c.Gender {
		case 0:
			return "male"
		case 1:
			return "female"
		}
	}

	return "n/a"
}

func (c *Character) GetFullJob() string {
	if c.JobName == nil {
		return "n/a"
	}

	job := *c.JobName

	if c.DepartmentName != nil {
		job += " / " + *c.DepartmentName
	} else {
		job += " / -"
	}

	if c.PositionName != nil {
		job += " / " + *c.PositionName
	} else {
		job += " / -"
	}

	return job
}

func (c *Character) String() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "\n=== Character: #%d\n", c.ID)

	fmt.Fprintf(&builder, " Name      : %s\n", c.GetFullName())

	if c.DateOfBirth != nil && *c.DateOfBirth != "" {
		fmt.Fprintf(&builder, " Birthday  : %s\n", *c.DateOfBirth)
	}

	fmt.Fprintf(&builder, " Gender    : %s\n", c.GetGender())

	if c.Playtime > 0 {
		fmt.Fprintf(&builder, " Playtime  : %d hours\n", c.Playtime/3600)
	}

	if c.LastSeen != nil && *c.LastSeen > 0 {
		fmt.Fprintf(&builder, " Last Seen : %s\n", time.Unix(*c.LastSeen, 0).Format(time.RFC822))
	}

	fmt.Fprintf(&builder, " Money     : $%d / $%d\n", c.Cash, c.Bank)
	fmt.Fprintf(&builder, " Job       : %s\n", c.GetFullJob())

	return builder.String()
}
