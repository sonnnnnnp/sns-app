package timeline

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	internal_errors "github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (uc *TimelineUsecase) GetUserTimeline(ctx context.Context, uID uuid.UUID, params *api.GetUserTimelineParams) (posts []api.Post, nextCursor uuid.UUID, err error) {
	selfUID := ctxhelper.GetUserID(ctx)

	queries := db.New(uc.pool)

	// ブロックされていないかを検証
	bs, err := queries.GetBlockStatus(ctx, db.GetBlockStatusParams{
		SelfID:   selfUID,
		TargetID: uID,
	})
	if err != nil {
		return nil, uuid.Nil, err
	}

	if bs.BlockedBy {
		return nil, uuid.Nil, internal_errors.ErrUserBlockedBy
	}

	// 指定されたカーソル投稿 ID から検索用日時を取得する
	var fromCursor *time.Time
	if params.Cursor != nil {
		r, err := queries.GetPostByID(ctx, db.GetPostByIDParams{
			SelfID: selfUID,
			PostID: *params.Cursor,
		})
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return nil, uuid.Nil, internal_errors.ErrCursorNotFound
			}
			return nil, uuid.Nil, err
		}
		fromCursor = &r.Post.CreatedAt.Time
	}

	// リミットの初期値と上限値を設定
	if params.Limit == nil {
		params.Limit = &defaultLimit
	}
	if *params.Limit > maxLimit {
		params.Limit = &maxLimit
	}

	// タイムラインを取得
	rows, err := queries.GetUserTimeline(ctx, db.GetUserTimelineParams{
		SelfID:    selfUID,
		AuthorID:  uID,
		CreatedAt: fromCursor,
		Limit:     int64(*params.Limit),
	})
	if err != nil {
		return nil, uuid.Nil, err
	}

	posts = make([]api.Post, 0)
	for _, r := range rows {
		posts = append(posts, api.Post{
			Author: api.User{
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

	// 最後の投稿 ID を次のカーソルとして返却する
	return posts, posts[len(posts)-1].Id, nil
}
