package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"dvdrental/graph/generated"
	"dvdrental/graph/model"
	"fmt"
)

// Actor is the resolver for the actor field.
func (r *queryResolver) Actor(ctx context.Context, id int) (*model.Actor, error) {
	panic(fmt.Errorf("not implemented"))
}

// Actors is the resolver for the actors field.
func (r *queryResolver) Actors(ctx context.Context, page *model.Pagination) ([]*model.Actor, error) {
	panic(fmt.Errorf("not implemented"))
}

// Film is the resolver for the film field.
func (r *queryResolver) Film(ctx context.Context, id int) (*model.Film, error) {
	film, err := r.service.GetFilm(ctx, id)
	if err != nil {
		return nil, err
	}
	return film.ToModel(), nil
}

// Films is the resolver for the films field.
func (r *queryResolver) Films(ctx context.Context, page *model.Pagination) ([]*model.Film, error) {
	offset := 0
	limit := 20

	if page != nil {
		if page.Offset != nil {
			offset = *page.Offset
		}

		if page.Limit != nil {
			limit = *page.Limit
		}
	}

	films, err := r.service.ListFilms(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	results := make([]*model.Film, len(films))
	for i, film := range films {
		results[i] = film.ToModel()
	}
	return results, nil
}

// Customer is the resolver for the customer field.
func (r *queryResolver) Customer(ctx context.Context, id int) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

// Customers is the resolver for the customers field.
func (r *queryResolver) Customers(ctx context.Context, page *model.Pagination) ([]*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
