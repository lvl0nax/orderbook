package main

import (
	"github.com/lvl0nax/orderbook/binance"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/orderbook/binance", binance.FetchOrderbook)

	app.Listen(":3000")
}
