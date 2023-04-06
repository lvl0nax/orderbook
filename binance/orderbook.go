package binance

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func FetchOrderbook(ctx *fiber.Ctx) error {
	symbol := ctx.Query("symbol")
	limit := ctx.Query("limit")
	if limit == "" {
		limit = "20"
	}

	url := fmt.Sprintf("https://api.binance.com/api/v3/depth?symbol=%s&limit=%s", symbol, limit)

	resp, err := http.Get(url)
	if err != nil {
		msg := fmt.Sprintln("Invalid request: ", url, " | ", err.Error())
		ctx.Status(fiber.StatusBadRequest).SendString(msg)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		msg := fmt.Sprintln("Something wrong with binance response: ", url, " | ", err.Error())
		ctx.Status(fiber.StatusBadRequest).SendString(msg)
	}

	return ctx.Send(body)
}
