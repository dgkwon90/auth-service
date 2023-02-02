package webapp

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func middleware(c *fiber.Ctx) error {
	fmt.Println("Don't mind me!")
	return c.Next()
}

func handler(c *fiber.Ctx) error {
	return c.SendString(c.Path() + " Hello, World 👋!")
}

func handlerList(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Hi John!",
	})
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handlerUser(c *fiber.Ctx) error {
	return c.JSON(&User{"John", 20})
}

func Start() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
		// => 404 "Not Found"
	})

	// Root API route
	api := app.Group("/api", middleware)

	v1 := api.Group("/v1", middleware)
	v1.Get("/list", handler)
	v1.Get("/user", handler)

	v2 := api.Group("/v2", middleware)
	v2.Get("/list", handlerList)
	v2.Get("/user", handlerUser)

	if err := app.Listen(":1090"); err != nil {
		log.Fatal("err: ", err.Error())
	}
}
