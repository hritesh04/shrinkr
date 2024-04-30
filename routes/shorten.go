package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/hritesh04/url-shortner/database"
)

type Request struct{
	Url string	
	CustomUrl string
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

	// check user plan using userId


	// set expiry according to the plan


	// set Rate remaining according to the plan


	if body.CustomUrl == "" {
		body.CustomUrl = uuid.New().String()[:6]
	}

	// add to db and send response

	return c.Status(200).JSON(fiber.Map{
		"success":true,
		"data":body,
	})
}