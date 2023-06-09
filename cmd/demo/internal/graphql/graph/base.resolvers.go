package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
	"github.com/go-nova/cmd/demo/internal/dao/model"
	"github.com/go-nova/cmd/demo/internal/graphql/generated"
	"github.com/go-nova/cmd/demo/module/demo/user_graphql/user_query"
)

// User is the resolver for the User field.
func (r *mutationResolver) User(ctx context.Context) (*model.UserMutationObject, error) {
	return &model.UserMutationObject{}, nil
}

// User is the resolver for the User field.
func (r *queryResolver) User(ctx context.Context) (*model.UserQueryObject, error) {
	return &model.UserQueryObject{}, nil
}

// Register is the resolver for the register field.
func (r *userMutationObjectResolver) Register(ctx context.Context, obj *model.UserMutationObject, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Register - register"))
}

// UserList is the resolver for the user_list field.
func (r *userQueryObjectResolver) UserList(ctx context.Context, obj *model.UserQueryObject, input *model.UserID) ([]*model.User, error) {
	var userQuery = &user_query.QueryResolver{}
	return userQuery.UserList(ctx, input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// UserMutationObject returns generated.UserMutationObjectResolver implementation.
func (r *Resolver) UserMutationObject() generated.UserMutationObjectResolver {
	return &userMutationObjectResolver{r}
}

// UserQueryObject returns generated.UserQueryObjectResolver implementation.
func (r *Resolver) UserQueryObject() generated.UserQueryObjectResolver {
	return &userQueryObjectResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userMutationObjectResolver struct{ *Resolver }
type userQueryObjectResolver struct{ *Resolver }
