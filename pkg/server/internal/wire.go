package internal

import (
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonnnnnnp/reverie/internal/adapter/controller"
	"github.com/sonnnnnnp/reverie/internal/usecase/authorize"
	"github.com/sonnnnnnp/reverie/internal/usecase/call"
	"github.com/sonnnnnnp/reverie/internal/usecase/call_timeline"
	"github.com/sonnnnnnp/reverie/internal/usecase/post"
	"github.com/sonnnnnnp/reverie/internal/usecase/stream"
	"github.com/sonnnnnnp/reverie/internal/usecase/timeline"
	"github.com/sonnnnnnp/reverie/internal/usecase/user"
	"github.com/sonnnnnnp/reverie/pkg/line"
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
