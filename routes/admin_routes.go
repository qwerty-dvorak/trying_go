package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qwerty-dvorak/trying_go/controller"
	"database/sql"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
    app.Get("/users", func(c *fiber.Ctx) error {
        users,err := controller.GetUsers(db)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
        return c.JSON(users)
    })
}