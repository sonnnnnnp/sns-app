package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonnnnnnp/reverie/internal/http/adapter/api"
)

type IUserUsecase interface {
	// users
	GetSelf(ctx context.Context) (*api.User, error)
	GetUser(ctx context.Context, params api.GetUserParams) (*api.User, error)

	UpdateUser(ctx context.Context, body *api.UpdateUserJSONBody) (*api.User, error)

	// blocks
	BlockUser(ctx context.Context, uID uuid.UUID) error

	GetUserBlocking(ctx context.Context) ([]api.User, error)

	UnblockUser(ctx context.Context, uID uuid.UUID) error

	// follows
	FollowUser(ctx context.Context, uID uuid.UUID) (*api.SocialConnection, error)

	GetUserFollowers(ctx context.Context, uID uuid.UUID) ([]api.UserFollower, error)
	GetUserFollowing(ctx context.Context, uID uuid.UUID) ([]api.UserFollower, error)

	UnfollowUser(ctx context.Context, uID uuid.UUID) (*api.SocialConnection, error)
	RemoveUserFromFollowers(ctx context.Context, uID uuid.UUID) (*api.SocialConnection, error)
}

type UserUsecase struct {
	pool *pgxpool.Pool
}

func New(
	pool *pgxpool.Pool,
) *UserUsecase {
	return &UserUsecase{
		pool: pool,
	}
}

var _ IUserUsecase = (*UserUsecase)(nil)
