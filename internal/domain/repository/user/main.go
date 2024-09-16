package user_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
)

type IUserRepository interface {
	GetUser(ctx context.Context, id uuid.UUID) (*ent.User, error)
	GetUserByLineID(ctx context.Context, lineID string) (*ent.User, error)
	CreateUser(ctx context.Context, id string) (*ent.User, error)
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
