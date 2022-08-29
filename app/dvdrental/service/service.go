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
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	actor, err := s.database.GetActor(ctx, int32(actorID))
	if err != nil {
		return nil, err
	}
	return entity.NewActor(actor), nil
}

func (s *service) ListActors(ctx context.Context, offset, limit int) ([]*entity.Actor, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	listActorsParams := &repository.ListActorsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	actors, err := s.database.ListActors(ctx, listActorsParams)
	if err != nil {
		return nil, err
	}

	entities := make([]*entity.Actor, len(actors))
	for i, actor := range actors {
		entities[i] = entity.NewActor(actor)
	}
	return entities, nil
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
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	listFilmsParams := &repository.ListFilmsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	films, err := s.database.ListFilms(ctx, listFilmsParams)
	if err != nil {
		return nil, err
	}

	entities := make([]*entity.Film, len(films))
	for i, film := range films {
		entities[i] = entity.NewFilm(film)
	}
	return entities, nil
}

func (s *service) GetCustomer(ctx context.Context, customerID int) (*entity.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	customer, err := s.database.GetCustomer(ctx, int32(customerID))
	if err != nil {
		return nil, err
	}
	return entity.NewCustomer(customer), nil
}

func (s *service) ListCustomers(ctx context.Context, offset, limit int) ([]*entity.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	listCustomersParams := &repository.ListCustomersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	customers, err := s.database.ListCustomers(ctx, listCustomersParams)
	if err != nil {
		return nil, err
	}

	entities := make([]*entity.Customer, len(customers))
	for i, customer := range customers {
		entities[i] = entity.NewCustomer(customer)
	}
	return entities, nil
}
