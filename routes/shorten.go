package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/hritesh04/url-shortner/database"
)

type Request struct{
	Url string
	customUrl string
	Expiry time.Duration
	UserId int32
}

func Shorten(c *fiber.Ctx)error{
	body:=Request{}
	if err := c.BodyParser(&body); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":false,
			"data":"can not parse JSON",
		})
	}

	db := database.Connect()
	defer db.Close()

	// rows,err := db.QueryRow("SELECT plan FROM users WHERE id = $1",body.UserId).Scan(&)

	if body.customUrl == "" {
		body.customUrl = uuid.New().String()[:6]
	}

	return c.Status(200).JSON(fiber.Map{
		"success":true,
		"data":body,
	})
}