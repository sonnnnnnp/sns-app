package internal

import (
	"github.com/google/wire"
	"github.com/sonnnnnnp/sns-app/internal/adapter/controller"
	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
	post_repository "github.com/sonnnnnnp/sns-app/internal/domain/repository/post"
	user_repository "github.com/sonnnnnnp/sns-app/internal/domain/repository/user"
	authorize_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/authorize"
	post_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/post"
	user_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/user"
	"github.com/sonnnnnnp/sns-app/pkg/line"
)

func Wire(db *ent.Client) *controller.Controller {
	wire.Build(
		line.New,

		post_repository.New,
		user_repository.New,

		authorize_usecase.New,
		post_usecase.New,
		user_usecase.New,

		controller.New,
	)

	return &controller.Controller{}
}
