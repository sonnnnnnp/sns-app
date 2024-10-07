package timeline

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/domain/post"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (tu *TimelineUsecase) GetTimeline(ctx context.Context, params *oapi.GetTimelineParams) (posts []*post.PostWithFavoriteCount, nextCursor uuid.UUID, err error) {
	posts, nextCursor, err = tu.postRepo.GetPosts(ctx, params.Limit, params.Cursor, params.UserId)
	if err != nil {
		return nil, uuid.Nil, err
	}

	return posts, nextCursor, nil
}
