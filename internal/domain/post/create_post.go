package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (repo *PostRepository) CreatePost(ctx context.Context, uID uuid.UUID, body *oapi.CreatePostJSONBody) (*ent.Post, error) {
	// query := pr.db.Post.Create().SetAuthorID(uID)

	// if body.Content != nil {
	// 	query.SetText(*body.Content)
	// }

	// sp, err := query.Save(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// return pr.db.Post.Query().Where(post.ID(sp.ID)).WithAuthor().Only(ctx)

	return nil, nil
}
