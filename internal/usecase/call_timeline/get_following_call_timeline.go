package call_timeline

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/reverie/internal/adapter/api"
	"github.com/sonnnnnnp/reverie/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/reverie/pkg/db"
)

func (uc *CallTimelineUsecase) GetFollowingCallTimeline(ctx context.Context, params api.GetFollowingCallTimelineParams) (rooms []api.CallRoom, nextCursor uuid.UUID, err error) {
	queries := db.New(uc.pool)

	// リミットの初期値と上限値を設定
	if params.Limit == nil {
		params.Limit = &defaultLimit
	}
	if *params.Limit > maxLimit {
		params.Limit = &maxLimit
	}

	rows, err := queries.GetFollowingCallTimeline(ctx, db.GetFollowingCallTimelineParams{
		SelfID: ctxhelper.GetUserID(ctx),
		Limit:  int64(*params.Limit),
	})
	if err != nil {
		return nil, uuid.Nil, err
	}

	rooms = make([]api.CallRoom, 0)
	for _, r := range rows {
		var dto api.CallRoom
		dto.Id = r.Call.ID
		dto.Title = *r.Call.Title
		dto.IsActive = !r.Call.EndedAt.Valid
		dto.JoinableBy = api.CallRoomJoinableBy(r.Call.JoinableBy)
		dto.Type = api.CallRoomType(r.Call.Type)

		var cp []api.CallParticipant
		err = json.Unmarshal(r.Participants, &cp)
		if err != nil {
			return nil, uuid.Nil, err
		}

		dto.Participants = cp

		rooms = append(rooms, dto)
	}

	if len(rooms) == 0 {
		return rooms, uuid.Nil, nil
	}

	// 最後のルーム ID を次のカーソルとして返却する
	return rooms, rooms[len(rooms)-1].Id, nil
}
