package post

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

var (
	defaultLimit = 25
	maxLimit     = 100
)

func (repo *PostRepository) GetTimeline(ctx context.Context, limit *int, fromCursor *time.Time, targetUID *uuid.UUID) (rows []db.GetPostsRow, err error) {
	uID := ctxhelper.GetUserID(ctx)

	if limit == nil {
		limit = &defaultLimit
	}
	if *limit > maxLimit {
		limit = &maxLimit
	}

	rows, err = db.New(repo.pool).GetPosts(ctx, db.GetPostsParams{
		UserID:    &uID,
		AuthorID:  targetUID,
		CreatedAt: fromCursor,
		Limit:     int64(*limit),
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}
