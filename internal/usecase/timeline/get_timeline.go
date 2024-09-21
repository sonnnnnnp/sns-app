package timeline_usecase

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
)

func (tu *TimelineUsecase) GetTimeline(ctx context.Context) ([]*ent.Post, error) {
	posts, err := tu.postRepo.GetPosts(ctx)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
