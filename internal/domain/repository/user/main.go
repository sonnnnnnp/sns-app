package user_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IUserRepository interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*ent.User, error)
	GetUserByLineID(ctx context.Context, lineID string) (*ent.User, error)
	CreateUser(ctx context.Context) (*ent.User, error)
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
