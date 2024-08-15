package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func Init() {
	dsn := "host=db user=user password=password dbname=db port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
	fmt.Println("successfully migrated")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		db.Create(&User{
			Name:  "gorm",
			Email: "gorm@example.com",
		})

		var user User
		db.Last(&user)

		return c.String(http.StatusOK, fmt.Sprintf("User ID: %d", user.ID))
	})

	e.Logger.Fatal(e.Start(":1323"))
}
