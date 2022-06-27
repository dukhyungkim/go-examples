package database

import (
	"fmt"
	"graphql/config"
	"graphql/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository interface {
	SaveLink(link *entity.Link) (int64, error)
	FetchLinks() ([]*entity.Link, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepository(cfg *config.RDB) (Repository, error) {
	const dsnTemplate = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s"
	dsn := fmt.Sprintf(dsnTemplate, cfg.Host, cfg.Username, cfg.Password, cfg.Database, cfg.Port, cfg.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &repo{db: db}, nil
}

func (r *repo) SaveLink(link *entity.Link) (int64, error) {
	if err := r.db.Create(link).Error; err != nil {
		return 0, err
	}
	return link.ID, nil
}

func (r *repo) FetchLinks() ([]*entity.Link, error) {
	var links []*entity.Link
	if err := r.db.Find(&links).Error; err != nil {
		return nil, err
	}
	return links, nil
}

func (r *repo) SaveUser(user *entity.User) (int64, error) {
	if err := r.db.Create(user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}
