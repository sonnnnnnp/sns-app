package repository

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, id string) (*ent.User, error)
}

type UserRepository struct {
	db *ent.Client
}

func NewUserRepository(db *ent.Client) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

var _ IUserRepository = (*UserRepository)(nil)
