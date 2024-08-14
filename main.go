package main

import(
	"github.com/qwerty-dvorak/trying_go/database"
	"github.com/qwerty-dvorak/trying_go/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := database.NewSession()
	print("\n")
	if err != nil {
		print(err)
	}
	print(db,"\n")
    println("Hello, Go!")
	app:=fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	routes.SetupRoutes(app, db)	
	app.Listen(":3000")
}