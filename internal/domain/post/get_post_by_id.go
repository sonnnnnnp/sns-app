package post

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *PostRepository) GetPostByID(ctx context.Context, pID uuid.UUID) (*db.GetPostByIDRow, error) {
	uID := ctxhelper.GetUserID(ctx)

	r, err := db.New(repo.pool).GetPostByID(ctx, db.GetPostByIDParams{
		SelfID: &uID,
		PostID: pID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &r, nil
}
