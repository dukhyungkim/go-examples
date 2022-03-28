package repository

import (
	"fmt"
	"go-examples/common/config"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMariaDB(rdb *config.RDB) (*Storage, error) {
	const dsnTemplate = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s"
	dsn := fmt.Sprintf(dsnTemplate, rdb.Username, rdb.Password, rdb.Host, rdb.Port, rdb.Database, url.QueryEscape(rdb.TimeZone))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func NewPostgreSQL(rdb *config.RDB) (*Storage, error) {
	const dsnTemplate = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s"
	dsn := fmt.Sprintf(dsnTemplate, rdb.Host, rdb.Username, rdb.Password, rdb.Database, rdb.Port, rdb.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}
