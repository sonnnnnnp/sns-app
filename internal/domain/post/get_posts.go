package post

import (
	"context"

	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/ent/post"
)

func (pr *PostRepository) GetPosts(ctx context.Context) ([]*ent.Post, error) {
	posts, err := pr.db.Post.Query().WithAuthor().Order(ent.Desc(post.FieldCreatedAt)).All(ctx)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
