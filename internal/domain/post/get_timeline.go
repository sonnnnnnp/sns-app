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

func (repo *PostRepository) GetTimeline(ctx context.Context, limit *int, fromCursor *uuid.UUID, uID *uuid.UUID) (rows []db.GetPostsRow, nextCursor uuid.UUID, err error) {
	queries := db.New(repo.pool)

	if limit == nil {
		limit = &defaultLimit
	}
	if *limit > maxLimit {
		limit = &maxLimit
	}

	var cursorTime *time.Time
	if fromCursor != nil {
		cursorPost, err := queries.GetPostByID(ctx, *fromCursor)
		if err != nil {
			return nil, uuid.Nil, err
		}
		cursorTime = &cursorPost.CreatedAt.Time
	}

	rows, err = queries.GetPosts(ctx, db.GetPostsParams{
		UserID:    ctxhelper.GetUserID(ctx),
		AuthorID:  uID,
		CreatedAt: cursorTime,
		Limit:     int64(*limit),
	})
	if err != nil {
		return nil, uuid.Nil, err
	}

	if len(rows) == 0 {
		return nil, uuid.Nil, nil
	}

	nextCursor = rows[len(rows)-1].Post.ID

	return rows, nextCursor, nil
}
