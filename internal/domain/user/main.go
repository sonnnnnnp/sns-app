package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonnnnnnp/sns-app/pkg/db"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, lineID *string) (*db.User, error)
	FollowUser(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) error
	GetSocialContext(ctx context.Context, selfID uuid.UUID, targetUID uuid.UUID) (*oapi.SocialContext, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*ent.User, error)
	GetUserByLineID(ctx context.Context, lineID string) (*db.User, error)
	GetUserByName(ctx context.Context, name string) (*ent.User, error)
	GetUserFollowers(ctx context.Context, uID uuid.UUID) ([]*ent.User, error)
	GetUserFollowing(ctx context.Context, uID uuid.UUID) ([]*ent.User, error)
	UnfollowUser(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) error
	UpdateUser(ctx context.Context, id uuid.UUID, data *oapi.UpdateUserJSONBody) (*ent.User, error)
}

type UserRepository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}

var _ IUserRepository = (*UserRepository)(nil)
