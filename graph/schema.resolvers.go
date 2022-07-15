package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/someday-94/TypeGoMongo-Server/graph/generated"
	"github.com/someday-94/TypeGoMongo-Server/model"
	"github.com/someday-94/TypeGoMongo-Server/repository"
)

var (
	memoRepo repository.MongoRepo = repository.NewMongoRepo()
)

// CreateMemo is the resolver for the createMemo field.
func (r *mutationResolver) CreateMemo(ctx context.Context, input model.NewMemo) (*model.Memo, error) {
	user := &model.User{
		ID:   input.UserId,
		Name: "user-" + input.UserId,
	}
	memo := &model.Memo{
		ID:      strconv.Itoa(rand.Int()),
		Content: input.Content,
		Author:  user,
	}

	memoRepo.Save(memo)

	return memo, nil
}

// Memos is the resolver for the memos field.
func (r *queryResolver) Memos(ctx context.Context) ([]*model.Memo, error) {
	return memoRepo.FindAll(), nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// MemoAdded is the resolver for the memoAdded field.
func (r *subscriptionResolver) MemoAdded(ctx context.Context, repoFullName string) (<-chan *model.Memo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
