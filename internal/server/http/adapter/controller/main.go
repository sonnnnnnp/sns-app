package controller

import (
	"github.com/sonnnnnnp/reverie/internal/server/http/adapter/api"
	"github.com/sonnnnnnp/reverie/internal/server/http/usecase/authorize"
	"github.com/sonnnnnnp/reverie/internal/server/http/usecase/call"
	"github.com/sonnnnnnp/reverie/internal/server/http/usecase/call_timeline"
	"github.com/sonnnnnnp/reverie/internal/server/http/usecase/gateway"
	"github.com/sonnnnnnp/reverie/internal/server/http/usecase/post"
	"github.com/sonnnnnnp/reverie/internal/server/http/usecase/timeline"
	"github.com/sonnnnnnp/reverie/internal/server/http/usecase/user"
)

type Controller struct {
	authUsecase         *authorize.AuthorizeUsecase
	callUsecase         *call.CallUsecase
	callTimelineUsecase *call_timeline.CallTimelineUsecase
	postUsecase         *post.PostUsecase
	gatewayUsecase      *gateway.GatewayUsecase
	timelineUsecase     *timeline.TimelineUsecase
	userUsecase         *user.UserUsecase
}

func New(
	authUsecase *authorize.AuthorizeUsecase,
	callUsecase *call.CallUsecase,
	callTimelineUsecase *call_timeline.CallTimelineUsecase,
	postUsecase *post.PostUsecase,
	gatewayUsecase *gateway.GatewayUsecase,
	timelineUsecase *timeline.TimelineUsecase,
	userUsecase *user.UserUsecase,
) *Controller {
	return &Controller{
		authUsecase:         authUsecase,
		callUsecase:         callUsecase,
		callTimelineUsecase: callTimelineUsecase,
		postUsecase:         postUsecase,
		gatewayUsecase:      gatewayUsecase,
		timelineUsecase:     timelineUsecase,
		userUsecase:         userUsecase,
	}
}

var _ api.ServerInterface = (*Controller)(nil)
