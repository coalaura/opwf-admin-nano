package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	db *gorm.DB

	ID       int       `json:"serviceId"`
	Name     *nullable `json:"serverName"`
	Address  string    `json:"address"`
	Port     int       `json:"port"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Database string    `json:"database"`
}

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

func (d *Database) Connect() error {
	db, err := gorm.Open(mysql.Open(d.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}

	d.db = db

	return nil
}

func (d *Database) FindUserByID(id string) (*User, error) {
	var user User

	res := d.db.Table("users").Take(&user, "user_id = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (d *Database) FindUserByLicense(license string) (*User, error) {
	var user User

	res := d.db.Table("users").Take(&user, "license_identifier = ?", license)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (d *Database) FindCharactersByID(id string) ([]Character, error) {
	var characters []Character

	res := d.db.Table("characters").Find(&characters, "character_id = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}

	return characters, nil
}

func (d *Database) FindCharactersByLicense(license string) ([]Character, error) {
	var characters []Character

	res := d.db.Table("characters").Find(&characters, "license_identifier = ?", license)
	if res.Error != nil {
		return nil, res.Error
	}

	return characters, nil
}
