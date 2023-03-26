package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/recepdmr/apigateways/public/pkg/post"
)

func main() {
	port := os.Getenv("PORT")

	app := fiber.New()

	post.RegisterPostHandler(app)
	app.Listen(fmt.Sprintf(":%s", port))
}
