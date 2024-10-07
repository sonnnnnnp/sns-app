package timeline

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/domain/post"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type ITimelineUsecase interface {
	GetTimeline(ctx context.Context, params *oapi.GetTimelineParams) (posts []*post.PostWithFavoriteCount, nextCursor uuid.UUID, err error)
}

type TimelineUsecase struct {
	postRepo *post.PostRepository
}

func New(
	postRepo *post.PostRepository,
) *TimelineUsecase {
	return &TimelineUsecase{
		postRepo: postRepo,
	}
}

var _ ITimelineUsecase = (*TimelineUsecase)(nil)
