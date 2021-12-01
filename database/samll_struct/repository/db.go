package repository

import (
	"fmt"
	"go-examples/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Storage struct {
	db *gorm.DB
}

func NewMariaDB(rdb *config.RDB) *Storage {
	const dsnTemplate = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(dsnTemplate, rdb.Username, rdb.Password, rdb.Host, rdb.Port, rdb.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	//if err := db.AutoMigrate(&entity.Person{}); err != nil {
	//	log.Fatalln(err)
	//}

	return &Storage{db: db}
}
