package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (c Controller) GetSelf(ctx echo.Context) error {
	uID := ctxhelper.GetUserID(ctx.Request().Context())

	u, err := c.userUsecase.GetUserByID(ctx.Request().Context(), uID)
	if err != nil {
		return err
	}

	if u == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	return c.json(ctx, http.StatusOK, &oapi.User{
		Id:             u.ID,
		Name:           u.Name,
		Nickname:       u.Nickname,
		AvatarImageUrl: &u.AvatarImageURL,
		BannerImageUrl: &u.BannerImageURL,
		Biography:      &u.Biography,
		UpdatedAt:      u.UpdatedAt,
		CreatedAt:      u.CreatedAt,
	})
}

func (c Controller) GetUserByName(ctx echo.Context, name string) error {
	u, sc, err := c.userUsecase.GetUserByName(ctx.Request().Context(), name)
	if err != nil {
		return err
	}

	if u == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	return c.json(ctx, http.StatusOK, &oapi.User{
		Id:             u.ID,
		Name:           u.Name,
		Nickname:       u.Nickname,
		AvatarImageUrl: &u.AvatarImageURL,
		BannerImageUrl: &u.BannerImageURL,
		Biography:      &u.Biography,
		SocialContext:  sc,
		UpdatedAt:      u.UpdatedAt,
		CreatedAt:      u.CreatedAt,
	})
}

func (c Controller) UpdateUser(ctx echo.Context) error {
	uID := ctxhelper.GetUserID(ctx.Request().Context())

	var body oapi.UpdateUserJSONBody
	if err := ctx.Bind(&body); err != nil {
		// return errors.ValidationFailed.Wrap(errors.WithStack(err), err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	u, err := c.userUsecase.UpdateUser(ctx.Request().Context(), uID, &body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to update a user")
	}

	return c.json(ctx, http.StatusOK, &oapi.User{
		Id:             u.Id,
		Name:           u.Name,
		Nickname:       u.Nickname,
		AvatarImageUrl: u.AvatarImageUrl,
		BannerImageUrl: u.BannerImageUrl,
		Biography:      u.Biography,
		UpdatedAt:      u.UpdatedAt,
		CreatedAt:      u.CreatedAt,
	})
}

func (c Controller) GetUserFollowing(ctx echo.Context, id uuid.UUID) error {
	users, err := c.userUsecase.GetUserFollowing(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	oapiUsers := make([]oapi.User, len(users))
	for i, u := range users {
		oapiUsers[i] = oapi.User{
			AvatarImageUrl: &u.AvatarImageURL,
			Biography:      &u.Biography,
			BannerImageUrl: &u.BannerImageURL,
			CreatedAt:      u.CreatedAt,
			Nickname:       u.Nickname,
			Id:             u.ID,
			UpdatedAt:      u.UpdatedAt,
			Name:           u.Name,
		}
	}

	return c.json(ctx, http.StatusOK, &oapi.Users{
		Users: oapiUsers,
	})
}

func (c Controller) GetUserFollowers(ctx echo.Context, id uuid.UUID) error {
	users, err := c.userUsecase.GetUserFollowers(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	oapiUsers := make([]oapi.User, len(users))
	for i, u := range users {
		oapiUsers[i] = oapi.User{
			AvatarImageUrl: &u.AvatarImageURL,
			Biography:      &u.Biography,
			BannerImageUrl: &u.BannerImageURL,
			CreatedAt:      u.CreatedAt,
			Nickname:       u.Nickname,
			Id:             u.ID,
			UpdatedAt:      u.UpdatedAt,
			Name:           u.Name,
		}
	}

	return c.json(ctx, http.StatusOK, &oapi.Users{
		Users: oapiUsers,
	})
}

func (c Controller) FollowUser(ctx echo.Context, id uuid.UUID) error {
	sc, err := c.userUsecase.FollowUser(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, sc)
}

func (c Controller) UnfollowUser(ctx echo.Context, id uuid.UUID) error {
	sc, err := c.userUsecase.UnfollowUser(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, sc)
}

func (c Controller) RemoveUserFromFollowers(ctx echo.Context, id uuid.UUID) error {
	sc, err := c.userUsecase.RemoveUserFromFollowers(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.json(ctx, http.StatusOK, sc)
}
