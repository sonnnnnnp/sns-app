package controller

import (
	"github.com/sonnnnnnp/reverie/server/adapter/http/api"
	"github.com/sonnnnnnp/reverie/server/usecase/authorize"
	"github.com/sonnnnnnp/reverie/server/usecase/call"
	"github.com/sonnnnnnp/reverie/server/usecase/call_timeline"
	"github.com/sonnnnnnp/reverie/server/usecase/gateway"
	"github.com/sonnnnnnp/reverie/server/usecase/post"
	"github.com/sonnnnnnp/reverie/server/usecase/timeline"
	"github.com/sonnnnnnp/reverie/server/usecase/user"
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
