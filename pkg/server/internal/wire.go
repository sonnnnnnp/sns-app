package internal

import (
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonnnnnnp/sns-app/internal/adapter/controller"
	authorize_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/authorize"
	post_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/post"
	stream_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/stream"
	user_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/user"
	"github.com/sonnnnnnp/sns-app/pkg/line"
)

func Wire(pool *pgxpool.Pool) *controller.Controller {
	wire.Build(
		line.New,

		authorize_usecase.New,
		post_usecase.New,
		stream_usecase.New,
		user_usecase.New,

		controller.New,
	)

	return &controller.Controller{}
}
