package controller

import (
	"github.com/sonnnnnnp/reverie/server/adapter/http/api"
	"github.com/sonnnnnnp/reverie/server/usecase/http/authorize"
	"github.com/sonnnnnnp/reverie/server/usecase/http/call"
	"github.com/sonnnnnnp/reverie/server/usecase/http/call_timeline"
	"github.com/sonnnnnnp/reverie/server/usecase/http/gateway"
	"github.com/sonnnnnnp/reverie/server/usecase/http/post"
	"github.com/sonnnnnnp/reverie/server/usecase/http/timeline"
	"github.com/sonnnnnnp/reverie/server/usecase/http/user"
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
