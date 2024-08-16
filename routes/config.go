package routes

import (
	"github.com/gofiber/fiber/v2"
	"database/sql"
)

func Setup(app *fiber.App, db *sql.DB){
	GetUsers(app, db)
	GetUser(app, db)
	AddUser(app, db)
	Signup(app, db)
}