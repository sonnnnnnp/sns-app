package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonnnnnnp/sns-app/pkg/db"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IUserRepository interface {
	// users
	CreateUser(ctx context.Context, lineID *string) (*db.User, error)

	GetUserByID(ctx context.Context, id uuid.UUID) (*db.User, error)
	GetUserByLineID(ctx context.Context, lineID string) (*db.User, error)
	GetUserByName(ctx context.Context, name string) (*db.User, error)

	UpdateUser(ctx context.Context, id uuid.UUID, data *oapi.UpdateUserJSONBody) error

	// blocks
	BlockUser(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) error

	GetUserBlocking(ctx context.Context, uID uuid.UUID) ([]db.GetUserBlockingRow, error)
	GetBlockStatus(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) (db.GetBlockStatusRow, error)

	UnblockUser(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) error

	// follows
	FollowUser(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) error

	GetUserFollowers(ctx context.Context, uID uuid.UUID) ([]db.GetUserFollowersRow, error)
	GetUserFollowing(ctx context.Context, uID uuid.UUID) ([]db.GetUserFollowingRow, error)
	GetSocialConnection(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) (*oapi.SocialConnection, error)

	UnfollowUser(ctx context.Context, selfUID uuid.UUID, targetUID uuid.UUID) error
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
