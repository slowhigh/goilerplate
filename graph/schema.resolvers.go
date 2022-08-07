package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/someday-94/TypeGoMongo-Server/graph/generated"
	"github.com/someday-94/TypeGoMongo-Server/model"
	"github.com/someday-94/TypeGoMongo-Server/utils"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, user model.UserInput) (*model.User, error) {
	defer utils.TryCatch()

	newUser := &model.User{
		ID:   strconv.FormatInt(time.Now().UnixNano(), 16),
		Name: user.Name,
	}

	r.userRepo.InsertOne(newUser)

	return newUser, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, user model.UserInput) (*model.User, error) {
	defer utils.TryCatch()

	updatedUser := &model.User{
		ID:   id,
		Name: user.Name,
	}

	r.userRepo.UpdateByID(id, updatedUser)

	return updatedUser, nil
}

// DeleleUser is the resolver for the deleleUser field.
func (r *mutationResolver) DeleleUser(ctx context.Context, id string) (*model.User, error) {
	defer utils.TryCatch()

	return r.userRepo.DeleteById(id), nil
}

// CreateMemo is the resolver for the createMemo field.
func (r *mutationResolver) CreateMemo(ctx context.Context, memo model.MemoInput) (*model.Memo, error) {
	defer utils.TryCatch()

	targetUser := r.userRepo.FindOneByName(memo.Author.Name)
	newMemo := &model.Memo{
		ID:      strconv.FormatInt(time.Now().UnixNano(), 16),
		Content: memo.Content,
		Author:  targetUser,
	}

	r.memoRepo.InsertOne(newMemo)

	return newMemo, nil
}

// UpdateMemo is the resolver for the updateMemo field.
func (r *mutationResolver) UpdateMemo(ctx context.Context, id string, memo model.MemoInput) (*model.Memo, error) {
	defer utils.TryCatch()

	targetUser := r.userRepo.FindOneByName(memo.Author.Name)
	updatedMemo := &model.Memo{
		ID:      id,
		Content: memo.Content,
		Author:  targetUser,
	}

	r.memoRepo.UpdateByID(id, updatedMemo)

	return updatedMemo, nil
}

// DeleteMemo is the resolver for the deleteMemo field.
func (r *mutationResolver) DeleteMemo(ctx context.Context, id string) (*model.Memo, error) {
	defer utils.TryCatch()

	return r.memoRepo.DeleteById(id), nil
}

// Memo is the resolver for the memo field.
func (r *queryResolver) Memo(ctx context.Context, id string) (*model.Memo, error) {
	defer utils.TryCatch()

	return r.memoRepo.FindOneById(id), nil
}

// Memos is the resolver for the memos field.
func (r *queryResolver) Memos(ctx context.Context) ([]*model.Memo, error) {
	return r.memoRepo.FindAll(), nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	defer utils.TryCatch()

	return r.userRepo.FindOneById(id), nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.userRepo.FindAll(), nil
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
