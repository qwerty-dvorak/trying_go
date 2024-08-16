package main

import(
	"github.com/qwerty-dvorak/trying_go/database"
	"github.com/qwerty-dvorak/trying_go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db, err := database.NewSession()
	if err != nil {
		print(err)
	}
	print(db,"\n")
    println("Hello, Go!")
	app:=fiber.New()

	app.Use(cors.New(cors.Config{
        AllowOrigins: "*", 
        AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.Setup(app, db)
	app.Listen(":3000")
}