//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
	"github.com/sonnnnnnp/sns-app/internal/adapter/controller"
	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
	user_repository "github.com/sonnnnnnp/sns-app/internal/domain/repository/user"
	user_usecase "github.com/sonnnnnnp/sns-app/internal/usecase/user"
)

func Init(db *ent.Client) *controller.Controller {
	wire.Build(
		user_repository.New,
		user_usecase.New,
		controller.New,
	)

	return &controller.Controller{}
}
