package routes

import (
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
	"github.com/hritesh04/url-shortner/models"
	qrcode "github.com/skip2/go-qrcode"
)


func GenerateQr(c *fiber.Ctx)error {
	body:=models.Request{}
	if err := c.BodyParser(&body); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":false,
			"data":"can not parse JSON",
		})
	}
	var png []byte
	png,err := qrcode.Encode(body.Url,qrcode.Highest,256)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":false,
			"data":"Error creating qrcode",
		})
	}

	pngbase64 := base64.StdEncoding.EncodeToString(png)

	//`data:image/png;base64,${data.data}`

	return c.Status(200).JSON(fiber.Map{
		"success":true,
		"data":pngbase64,
	})
}