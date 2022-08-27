package service

import (
	"context"
	"dvdrental/graph/entity"
	"dvdrental/repository"
	"time"
)

type Service interface {
	GetActor(ctx context.Context, actorID int) (*entity.Actor, error)
	ListActors(ctx context.Context, offset, limit int) ([]*entity.Actor, error)

	GetFilm(ctx context.Context, filmID int) (*entity.Film, error)
	ListFilms(ctx context.Context, offset, limit int) ([]*entity.Film, error)

	GetCustomer(ctx context.Context, customerID int) (*entity.Customer, error)
	ListCustomers(ctx context.Context, offset, limit int) ([]*entity.Customer, error)
}

type service struct {
	database repository.Querier
}

func NewService(database repository.Querier) Service {
	return &service{
		database: database,
	}
}

func (s *service) GetActor(ctx context.Context, actorID int) (*entity.Actor, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) ListActors(ctx context.Context, offset, limit int) ([]*entity.Actor, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetFilm(ctx context.Context, filmID int) (*entity.Film, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	film, err := s.database.GetFilm(ctx, int32(filmID))
	if err != nil {
		return nil, err
	}
	return entity.NewFilm(film), nil
}

func (s *service) ListFilms(ctx context.Context, offset, limit int) ([]*entity.Film, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetCustomer(ctx context.Context, customerID int) (*entity.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) ListCustomers(ctx context.Context, offset, limit int) ([]*entity.Customer, error) {
	//TODO implement me
	panic("implement me")
}
