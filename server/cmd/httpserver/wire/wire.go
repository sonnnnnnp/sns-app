package commands

import (
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sonnnnnnp/reverie/server/adapter/http/controller"
	"github.com/sonnnnnnp/reverie/server/pkg/line"
	"github.com/sonnnnnnp/reverie/server/usecase/http/authorize"
	"github.com/sonnnnnnp/reverie/server/usecase/http/call"
	"github.com/sonnnnnnp/reverie/server/usecase/http/call_timeline"
	"github.com/sonnnnnnp/reverie/server/usecase/http/gateway"
	"github.com/sonnnnnnp/reverie/server/usecase/http/post"
	"github.com/sonnnnnnp/reverie/server/usecase/http/timeline"
	"github.com/sonnnnnnp/reverie/server/usecase/http/user"
)

func Wire(pool *pgxpool.Pool) *controller.Controller {
	wire.Build(
		line.New,

		authorize.New,
		call.New,
		call_timeline.New,
		post.New,
		gateway.New,
		timeline.New,
		user.New,

		controller.New,
	)

	return &controller.Controller{}
}
