// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: end_call.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const endCall = `-- name: EndCall :exec
UPDATE
    calls
SET
    ended_at = NOW()
WHERE
    calls.id = $1::uuid
`

func (q *Queries) EndCall(ctx context.Context, callID uuid.UUID) error {
	_, err := q.db.Exec(ctx, endCall, callID)
	return err
}
