package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (repo *PostRepository) CreatePost(ctx context.Context, uID uuid.UUID, body *api.CreatePostJSONBody) (pID uuid.UUID, err error) {
	queries := db.New(repo.pool)

	return queries.CreatePost(ctx, db.CreatePostParams{
		AuthorID: uID,
		Text:     body.Content,
	})
}
