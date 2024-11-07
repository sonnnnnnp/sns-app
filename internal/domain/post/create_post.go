package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/db"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (repo *PostRepository) CreatePost(ctx context.Context, uID uuid.UUID, body *oapi.CreatePostJSONBody) (pID uuid.UUID, err error) {
	queries := db.New(repo.pool)

	return queries.CreatePost(ctx, db.CreatePostParams{
		AuthorID: uID,
		Text:     body.Content,
	})
}
