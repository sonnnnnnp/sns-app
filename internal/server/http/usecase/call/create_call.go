package call

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/sonnnnnnp/reverie/internal/pkg/ctxhelper"
	internal_errors "github.com/sonnnnnnp/reverie/internal/pkg/errors"
	"github.com/sonnnnnnp/reverie/internal/server/http/adapter/api"
	"github.com/sonnnnnnp/reverie/pkg/db"
)

func (uc *CallUsecase) CreateCall(ctx context.Context, body api.CreateCallJSONBody) (*api.CallRoom, error) {
	// トランザクションを開始
	tx, err := uc.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	txQueries := db.New(uc.pool).WithTx(tx)

	selfUID := ctxhelper.GetUserID(ctx)

	// ユーザーが他の通話に参加していないか検証
	joined, err := txQueries.HasUserActiveCall(ctx, selfUID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	if joined {
		return nil, internal_errors.ErrAlreadyJoinedACall
	}

	// 通話を作成
	cID, err := txQueries.CreateCall(ctx, db.CreateCallParams{
		HostID:     selfUID,
		Title:      *body.Title,
		Type:       db.CallType(*body.Type),
		JoinableBy: db.CallJoinableBy(*body.JoinableBy),
	})
	if err != nil {
		return nil, err
	}

	// 通話のホストなので参加
	if err := txQueries.JoinCall(ctx, db.JoinCallParams{
		CallID:        cID,
		ParticipantID: selfUID,
		Role: db.NullCallParticipantRole{
			CallParticipantRole: db.CallParticipantRole(db.CallParticipantRoleHost),
			Valid:               true,
		},
	}); err != nil {
		return nil, err
	}

	// 作成した通話のデータを返却
	row, err := txQueries.GetCallByID(ctx, db.GetCallByIDParams{
		SelfID: selfUID,
		CallID: cID,
	})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	var dto api.CallRoom
	dto.Id = row.Call.ID
	dto.Title = *row.Call.Title
	dto.IsActive = !row.Call.EndedAt.Valid
	dto.JoinableBy = api.CallRoomJoinableBy(row.Call.JoinableBy)
	dto.Type = api.CallRoomType(row.Call.Type)

	var cp []api.CallParticipant
	err = json.Unmarshal(row.Participants, &cp)
	if err != nil {
		return nil, err
	}

	log.Printf("%v", cp)

	dto.Participants = cp

	return &dto, nil
}
