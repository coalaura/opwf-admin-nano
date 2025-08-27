package main

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	ID            int64
	PlayerName    *string
	DiscordId     *string
	IsTrusted     bool
	IsStaff       bool
	IsSeniorStaff bool
	IsSuperAdmin  bool
	Playtime      int64
	LastSeen      *int64
	AveragePing   *int64
	AverageFPS    *int64
	MediaDevices  string

	Characters []Character
}

type Character struct {
	ID             int64
	Created        bool
	CreatedAt      *int64
	FirstName      *string
	LastName       *string
	DateOfBirth    *string
	Gender         *int64
	Cash           int64
	Bank           int64
	LastSeen       *int64
	Playtime       int64
	JobName        *string
	DepartmentName *string
	PositionName   *string
}

func (u *User) GetRank() string {
	switch {
	case u.IsSuperAdmin:
		return "super-admin"
	case u.IsSeniorStaff:
		return "senior-staff"
	case u.IsStaff:
		return "staff"
	case u.IsTrusted:
		return "trusted"
	}

	return "player"
}

func (u *User) String() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "=== User: %d (%s)\n", u.ID, u.GetRank())

	if u.PlayerName != nil && *u.PlayerName != "" {
		fmt.Fprintf(&builder, " Player Name  : %s\n", *u.PlayerName)
	}

	if u.DiscordId != nil && *u.DiscordId != "" {
		fmt.Fprintf(&builder, " Discord ID   : %s\n", *u.DiscordId)
	}

	fmt.Fprintf(&builder, " Playtime     : %d hours\n", u.Playtime/3600)

	if u.LastSeen != nil && *u.LastSeen > 0 {
		fmt.Fprintf(&builder, " Last Seen    : %s\n", time.Unix(*u.LastSeen, 0).Format(time.RFC822))
	}

	if u.AveragePing != nil && *u.AveragePing > 0 {
		fmt.Fprintf(&builder, " Avg. Ping    : %d ms\n", *u.AveragePing)
	}

	if u.AverageFPS != nil && *u.AverageFPS > 0 {
		fmt.Fprintf(&builder, " Avg. FPS     : %d\n", *u.AverageFPS)
	}

	if u.MediaDevices != "" {
		fmt.Fprintf(&builder, " MediaDevices : %s\n", u.MediaDevices)
	}

	if len(u.Characters) > 0 {
		for _, character := range u.Characters {
			fmt.Fprintf(&builder, "\n=== Character: #%d\n", character.ID)

			builder.WriteString(character.String())
		}
	}

	return builder.String()
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
