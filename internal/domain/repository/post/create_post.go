package post_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (pr *PostRepository) CreatePost(ctx context.Context, uID uuid.UUID, body *oapi.CreatePostJSONBody) (*ent.Post, *ent.User, error) {
	query := pr.db.Post.Create().SetAuthorID(uID)

	if body.Content != nil {
		query.SetContent(*body.Content)
	}

	p, err := query.Save(ctx)
	if err != nil {
		return nil, nil, err
	}

	u, err := pr.db.User.Get(ctx, p.AuthorID)
	if err != nil {
		return nil, nil, err
	}

	return p, u, nil
}
