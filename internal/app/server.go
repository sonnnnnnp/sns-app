package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"

	"github.com/sonnnnnnp/sns-app/internal/pkg/db"
	"github.com/sonnnnnnp/sns-app/internal/pkg/oapi"
	"github.com/sonnnnnnp/sns-app/pkg/config"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

type UserController struct{}

func (u UserController) GetUsers(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, oapi.User{
		Id:    1,
		Name:  "admin",
		Email: "user@example.com",
	})
}

func Init(cfg *config.Config) {
	db, err := db.Open(db.Options{
		Host:     cfg.DBHost,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		Name:     cfg.DBName,
		Port:     cfg.DBPort,
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
	fmt.Println("successfully migrated")

	e := echo.New()

	uc := UserController{}

	oapi.RegisterHandlers(e, uc)

	// e.GET("/", func(c echo.Context) error {
	// 	db.Create(&User{
	// 		Name:  "gorm",
	// 		Email: "gorm@example.com",
	// 	})

	// 	var user User
	// 	db.Last(&user)

	// 	return c.String(http.StatusOK, fmt.Sprintf("User ID: %d", user.ID))
	// })

	e.Logger.Fatal(e.Start(":1323"))
}
