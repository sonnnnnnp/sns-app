package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/ent/post"
)

func (pr *PostRepository) GetPosts(ctx context.Context, limit *int, fromCursor *uuid.UUID) (posts []*ent.Post, nextCursor uuid.UUID, err error) {
	defaultLimit := 25
	if limit == nil {
		limit = &defaultLimit
	}

	maxLimit := 100
	if *limit > maxLimit {
		limit = &maxLimit
	}

	query := pr.db.Post.Query().WithAuthor().Order(ent.Desc(post.FieldCreatedAt)).Limit(*limit)

	if fromCursor != nil {
		cursorPost, err := pr.db.Post.Get(ctx, *fromCursor)
		if err != nil {
			return nil, uuid.Nil, err
		}

		query = query.Where(post.CreatedAtLT(cursorPost.CreatedAt))
	}

	posts, err = query.All(ctx)
	if err != nil {
		return nil, uuid.Nil, err
	}

	if len(posts) == 0 {
		return posts, uuid.Nil, nil
	}

	nextCursor = posts[len(posts)-1].ID

	return posts, nextCursor, nil
}
