package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qwerty-dvorak/trying_go/controller"
	"github.com/qwerty-dvorak/trying_go/schema"
	"database/sql"
)

func GetUsers(app *fiber.App, db *sql.DB) {
    app.Get("/users", func(c *fiber.Ctx) error {
        users,err := controller.ReadUsers(db)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
        return c.JSON(users)
    })
}

func GetUser(app *fiber.App, db *sql.DB) {
	app.Get("/users/:email", func(c *fiber.Ctx) error {
		email:=c.Params("email")
		user,err := controller.ReadUser(db, email)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(user)
	})
}

func AddUser(app *fiber.App, db *sql.DB) {
	app.Post("/users", func(c *fiber.Ctx) error {
		var user schema.CreateUser
		user.Email = c.Get("email")
		user.Password = c.Get("password")
		status,err := controller.CreateUser(db, user)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(status)
	})
}

func Signup(app *fiber.App, db *sql.DB) {
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("Login page")
	})	
}