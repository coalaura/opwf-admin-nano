package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type User struct {
	ID            int64   `gorm:"column:user_id;<-:false"`
	PlayerName    *string `gorm:"column:player_name;<-:false"`
	DiscordId     *string `gorm:"column:discord_id;<-:false"`
	IsTrusted     bool    `gorm:"column:is_trusted;<-:false"`
	IsStaff       bool    `gorm:"column:is_staff;<-:false"`
	IsSeniorStaff bool    `gorm:"column:is_senior_staff;<-:false"`
	IsSuperAdmin  bool    `gorm:"column:is_super_admin;<-:false"`
	Playtime      int64   `gorm:"column:playtime;<-:false"`
	LastSeen      *int64  `gorm:"column:last_seen;<-:false"`
	AveragePing   *int64  `gorm:"column:average_ping;<-:false"`
	AverageFPS    *int64  `gorm:"column:average_fps;<-:false"`
	MediaDevices  string  `gorm:"column:media_devices;<-:false"`
}

func HandleUser(database *Database, where string) error {
	var (
		user *User
		err  error
	)

	if IsValidInteger(where) {
		user, err = database.FindUserByID(where)
	} else if IsValidLicense(where) {
		user, err = database.FindUserByLicense(where)
	} else {
		err = errors.New("invalid where (allowed: id, license)")
	}

	if err != nil {
		return err
	}

	log.Println()
	log.Println(user)

	return nil
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

	return builder.String()
}
