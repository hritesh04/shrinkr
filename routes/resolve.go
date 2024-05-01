package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hritesh04/url-shortner/database"
	"github.com/hritesh04/url-shortner/models"
)



func Resolve(c *fiber.Ctx)error {
	url := c.Params("url")

	urlDetails:= models.Url{}

	db := database.Connect()
	defer db.Close()

	rows,err:=db.Query("SELECT original FROM urls WHERE shortened = $1",url)
	if err != nil{
		fmt.Println(err)
		return c.Status(fiber.StatusNoContent).JSON(&fiber.Map{
			"success":false,
			"data":"Invalid url",
		})
	}
	for rows.Next(){
		err := rows.Scan(&urlDetails.Original)
		if err != nil{
			fmt.Println(err)
			return c.Status(400).JSON(&fiber.Map{
				"success":false,
				"data":"Error scaning data",
			})
		}
	}

	return c.Redirect(urlDetails.Original,302)
}