package internal

import (
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonnnnnnp/sns-app/internal/adapter/controller"
	"github.com/sonnnnnnp/sns-app/internal/usecase/authorize"
	"github.com/sonnnnnnp/sns-app/internal/usecase/call"
	"github.com/sonnnnnnp/sns-app/internal/usecase/call_timeline"
	"github.com/sonnnnnnp/sns-app/internal/usecase/post"
	"github.com/sonnnnnnp/sns-app/internal/usecase/stream"
	"github.com/sonnnnnnp/sns-app/internal/usecase/timeline"
	"github.com/sonnnnnnp/sns-app/internal/usecase/user"
	"github.com/sonnnnnnp/sns-app/pkg/line"
)

func Wire(pool *pgxpool.Pool) *controller.Controller {
	wire.Build(
		line.New,

		authorize.New,
		call.New,
		call_timeline.New,
		post.New,
		stream.New,
		timeline.New,
		user.New,

		controller.New,
	)

	return &controller.Controller{}
}
