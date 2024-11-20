// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: delete_post.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const deletePost = `-- name: DeletePost :exec
DELETE FROM
    posts
WHERE
    id = $1::uuid
`

func (q *Queries) DeletePost(ctx context.Context, postID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deletePost, postID)
	return err
}
