package timeline

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (tu *TimelineUsecase) GetTimeline(ctx context.Context, params *oapi.GetTimelineParams) (posts []*ent.Post, nextCursor uuid.UUID, err error) {
	posts, nextCursor, err = tu.postRepo.GetPosts(ctx, params.Limit, params.Cursor)
	if err != nil {
		return nil, uuid.Nil, err
	}

	return posts, nextCursor, nil
}
