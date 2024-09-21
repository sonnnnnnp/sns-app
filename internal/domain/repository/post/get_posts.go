package post_repository

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
)

func (pr *PostRepository) GetPosts(ctx context.Context) ([]*ent.Post, error) {
	posts, err := pr.db.Post.Query().WithAuthor().All(ctx)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
