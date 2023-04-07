// +heroku goVersion go1.20.2

package main

import (
	"os"

	"github.com/lvl0nax/orderbook/binance"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/orderbook/binance", binance.FetchOrderbook)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen("0.0.0.0:" + port)
}
