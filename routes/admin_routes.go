package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qwerty-dvorak/trying_go/controller"
	"github.com/qwerty-dvorak/trying_go/schema"
	"database/sql"
)

func GetUsers(app *fiber.App, db *sql.DB) {
    app.Get("/users", func(c *fiber.Ctx) error {
        users := controller.ReadUsers(db)
		if users.Status != "success" {
			return c.Status(500).SendString(users.Message)
		}
        return c.JSON(users)
    })
}

func GetUser(app *fiber.App, db *sql.DB) {
	app.Get("/users/:email", func(c *fiber.Ctx) error {
		email:=c.Params("email")
		user:= controller.ReadUser(db, email)
		if user.Status != "success" {
			return c.Status(500).SendString(user.Message)
		}
		return c.JSON(user)
	})
}

func AddUser(app *fiber.App, db *sql.DB) {
	app.Post("/users", func(c *fiber.Ctx) error {
		var user schema.CreateUser
		user.Email = c.Get("email")
		user.Password = c.Get("password")
		newuser:= controller.CreateUser(db, user)
		if newuser.Status != "success" {
			return c.Status(500).SendString(newuser.Message)
		}
		return c.JSON(newuser)
	})
}

func Signup(app *fiber.App, db *sql.DB) {
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("Login page")
	})	
}