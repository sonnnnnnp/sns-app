package internal

import (
	"github.com/google/wire"
	"github.com/sonnnnnnp/sns-app/internal/adapter/controller"
	post_repository "github.com/sonnnnnnp/sns-app/internal/domain/post"
	user_repository "github.com/sonnnnnnp/sns-app/internal/domain/user"
	authorize_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/authorize"
	post_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/post"
	timeline_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/timeline"
	user_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/user"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/line"
)

func Wire(db *ent.Client) *controller.Controller {
	wire.Build(
		line.New,

		post_repository.New,
		user_repository.New,

		authorize_usecase.New,
		post_usecase.New,
		timeline_usecase.New,
		user_usecase.New,

		controller.New,
	)

	return &controller.Controller{}
}
