package call

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	internal_errors "github.com/sonnnnnnp/reverie/internal/errors"
	"github.com/sonnnnnnp/reverie/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/reverie/pkg/db"
)

func (uc *CallUsecase) EndCall(ctx context.Context, cID uuid.UUID) error {
	queries := db.New(uc.pool)

	selfUID := ctxhelper.GetUserID(ctx)

	row, err := queries.GetCallByID(ctx, db.GetCallByIDParams{
		SelfID: selfUID,
		CallID: cID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return internal_errors.ErrCallNotFound
		}
		return err
	}

	// 通話のホストか確認
	if row.Call.HostID != selfUID {
		return internal_errors.ErrNotCallHost
	}

	// 通話が既に終了しているか確認
	if row.Call.EndedAt.Valid {
		return internal_errors.ErrCallAlreadyEnded
	}

	return queries.EndCall(ctx, cID)
}
