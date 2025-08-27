package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func (d *Database) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		d.Username,
		d.Password,
		d.Address,
		d.Port,
		d.Database,
	)
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) Connect() error {
	db, err := sql.Open("mysql", d.DSN())
	if err != nil {
		return err
	}

	d.db = db

	return nil
}

func (d *Database) Find(license string) (*User, error) {
	var user User

	row := d.db.QueryRow("SELECT user_id, player_name, discord_id, is_trusted, is_staff, is_senior_staff, is_super_admin, playtime, average_ping, average_fps, last_seen, media_devices FROM users WHERE license_identifier = ? LIMIT 1", license)

	err := row.Scan(&user.ID, &user.PlayerName, &user.DiscordId, &user.IsTrusted, &user.IsStaff, &user.IsSeniorStaff, &user.IsSuperAdmin, &user.Playtime, &user.AveragePing, &user.AverageFPS, &user.LastSeen, &user.MediaDevices)
	if err != nil {
		return nil, err
	}

	rows, err := d.db.Query("SELECT character_id, character_created, character_creation_timestamp, first_name, last_name, date_of_birth, gender, cash, bank, job_name, department_name, position_name, playtime, last_seen FROM characters WHERE license_identifier = ? AND character_deleted != 1 ORDER BY character_id", license)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var character Character

		err := rows.Scan(
			&character.ID,
			&character.Created,
			&character.CreatedAt,
			&character.FirstName,
			&character.LastName,
			&character.DateOfBirth,
			&character.Gender,
			&character.Cash,
			&character.Bank,
			&character.JobName,
			&character.DepartmentName,
			&character.PositionName,
			&character.Playtime,
			&character.LastSeen,
		)

		if err != nil {
			return nil, err
		}

		user.Characters = append(user.Characters, character)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}
