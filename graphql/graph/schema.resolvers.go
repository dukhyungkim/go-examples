package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql/entity"
	"graphql/graph/generated"
	"graphql/graph/model"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var link model.Link
	link.Title = input.Title
	link.Address = input.Address

	linkEntity, err := entity.NewLink(&link)
	if err != nil {
		return nil, err
	}

	linkID, err := r.repo.SaveLink(linkEntity)
	if err != nil {
		return nil, err
	}

	link.ID = fmt.Sprint(linkID)
	return &link, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input *model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	fetchedLinks, err := r.repo.FetchLinks()
	if err != nil {
		return nil, err
	}

	resultLinks := make([]*model.Link, len(fetchedLinks))
	for i, link := range fetchedLinks {
		resultLinks[i] = link.ToModel()
	}
	return resultLinks, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
