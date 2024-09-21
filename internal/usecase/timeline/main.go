package timeline_usecase

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
	post_repository "github.com/sonnnnnnp/sns-app/internal/domain/repository/post"
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
