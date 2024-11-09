package post

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uc *PostUsecase) GetTimeline(ctx context.Context, params *oapi.GetTimelineParams) (posts []oapi.Post, nextCursor uuid.UUID, err error) {
	selfUID := ctxhelper.GetUserID(ctx)

	if params.UserId != nil {
		bs, err := uc.userRepo.GetBlockStatus(ctx, selfUID, *params.UserId)
		if err != nil {
			return nil, uuid.Nil, err
		}

		if bs.Blocking {
			return nil, uuid.Nil, errors.ErrUserBlockedByYou
		}
		if bs.BlockedBy {
			return nil, uuid.Nil, errors.ErrUserBlockingYou
		}
	}

	var fromCursor *time.Time
	if params.Cursor != nil {
		r, err := uc.postRepo.GetPostByID(ctx, *params.Cursor)
		if err != nil {
			return nil, uuid.Nil, err
		}
		if r == nil {
			return nil, uuid.Nil, errors.ErrCursorNotFound
		}
		fromCursor = &r.Post.CreatedAt.Time
	}

	rows, err := uc.postRepo.GetTimeline(ctx, params.Limit, fromCursor, params.UserId, params.Following)
	if err != nil {
		return nil, uuid.Nil, err
	}

	posts = make([]oapi.Post, 0)
	for _, r := range rows {
		posts = append(posts, oapi.Post{
			Author: oapi.User{
				AvatarImageUrl: r.User.AvatarImageUrl,
				BannerImageUrl: r.User.BannerImageUrl,
				Biography:      r.User.Biography,
				CreatedAt:      r.User.CreatedAt.Time,
				Id:             r.User.ID,
				Name:           r.User.Name,
				Nickname:       r.User.Nickname,
				UpdatedAt:      r.User.UpdatedAt.Time,
			},
			CreatedAt:      r.Post.CreatedAt.Time,
			Favorited:      r.Favorited,
			FavoritesCount: int(r.FavoritesCount),
			Id:             r.Post.ID,
			Text:           r.Post.Text,
			UpdatedAt:      r.Post.UpdatedAt.Time,
		})
	}

	if len(posts) == 0 {
		return posts, uuid.Nil, nil
	}

	return posts, posts[len(posts)-1].Id, nil
}
