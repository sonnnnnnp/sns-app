package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, lineID *string) (*ent.User, error)
	FollowUser(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) error
	GetSocialContext(ctx context.Context, selfID uuid.UUID, targetUID uuid.UUID) (*oapi.SocialContext, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*ent.User, error)
	GetUserByLineID(ctx context.Context, lineID string) (*ent.User, error)
	GetUserByName(ctx context.Context, name string) (*ent.User, error)
	GetUserFollowers(ctx context.Context, uID uuid.UUID) ([]*ent.User, error)
	GetUserFollowing(ctx context.Context, uID uuid.UUID) ([]*ent.User, error)
	UnfollowUser(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) error
	UpdateUser(ctx context.Context, id uuid.UUID, data *oapi.UpdateUserJSONBody) (*ent.User, error)
}

type UserRepository struct {
	db *ent.Client
}

func New(db *ent.Client) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

var _ IUserRepository = (*UserRepository)(nil)
