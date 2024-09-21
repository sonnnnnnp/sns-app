package post_repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
	"github.com/sonnnnnnp/sns-app/internal/domain/ent/post"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (pr *PostRepository) CreatePost(ctx context.Context, uID uuid.UUID, body *oapi.CreatePostJSONBody) (*ent.Post, error) {
	query := pr.db.Post.Create().SetAuthorID(uID)

	if body.Content != nil {
		query.SetContent(*body.Content)
	}

	sp, err := query.Save(ctx)
	if err != nil {
		return nil, err
	}

	p, err := pr.db.Post.Query().Where(post.IDEQ(sp.ID)).WithAuthor().Only(ctx)
	if err != nil {
		return nil, err
	}

	return p, nil
}
