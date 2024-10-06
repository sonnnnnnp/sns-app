package timeline

import (
	"context"

	"github.com/google/uuid"
	post_repository "github.com/sonnnnnnp/sns-app/internal/domain/post"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type ITimelineUsecase interface {
	GetTimeline(ctx context.Context, params *oapi.GetTimelineParams) (posts []*ent.Post, nextCursor uuid.UUID, err error)
}

type TimelineUsecase struct {
	postRepo *post_repository.PostRepository
}

func New(
	postRepo *post_repository.PostRepository,
) *TimelineUsecase {
	return &TimelineUsecase{
		postRepo: postRepo,
	}
}

var _ ITimelineUsecase = (*TimelineUsecase)(nil)
