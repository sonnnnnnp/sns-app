package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/reverie/server/adapter/http/api"
)

// users

func (c Controller) SearchUsersTypeahead(ctx echo.Context, params api.SearchUsersTypeaheadParams) error {
	// TODO: 最小限のデータを含むユーザーリストを custom_id で検索して返却する
	return nil
}
