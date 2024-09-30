package timeline

import (
	"context"

	post_repository "github.com/sonnnnnnp/sns-app/internal/domain/post"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
)

type ITimelineUsecase interface {
	GetTimeline(ctx context.Context) ([]*ent.Post, error)
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
