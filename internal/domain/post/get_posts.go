package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/ent/favorite"
	"github.com/sonnnnnnp/sns-app/pkg/ent/post"
	"github.com/sonnnnnnp/sns-app/pkg/ent/user"
)

type PostWithFavoriteCount struct {
	*ent.Post
	FavoriteCount int `json:"favorite_count"`
}

func (pr *PostRepository) GetPosts(ctx context.Context, limit *int, fromCursor *uuid.UUID, uID *uuid.UUID) (posts []*PostWithFavoriteCount, nextCursor uuid.UUID, err error) {
	defaultLimit := 25
	if limit == nil {
		limit = &defaultLimit
	}

	maxLimit := 100
	if *limit > maxLimit {
		limit = &maxLimit
	}

	query := pr.db.Post.Query().WithAuthor().Order(ent.Desc(post.FieldCreatedAt)).Limit(*limit)

	if uID != nil {
		query = query.Where(post.HasAuthorWith(user.ID(*uID)))
	}

	if fromCursor != nil {
		cursorPost, err := pr.db.Post.Get(ctx, *fromCursor)
		if err != nil {
			return nil, uuid.Nil, err
		}

		query = query.Where(post.CreatedAtLT(cursorPost.CreatedAt))
	}

	entPosts, err := query.All(ctx)
	if err != nil {
		return nil, uuid.Nil, err
	}

	if len(entPosts) == 0 {
		return nil, uuid.Nil, nil
	}

	for _, p := range entPosts {
		count, err := pr.db.Favorite.Query().Where(favorite.HasPostWith(post.ID(p.ID))).Count(ctx)
		if err != nil {
			return nil, uuid.Nil, err
		}

		posts = append(posts, &PostWithFavoriteCount{
			Post:          p,
			FavoriteCount: count,
		})
	}

	nextCursor = entPosts[len(entPosts)-1].ID

	return posts, nextCursor, nil
}
