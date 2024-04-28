package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hritesh04/url-shortner/database"
)



func Resolve(c *fiber.Ctx)error {
	url := c.Params("url")

	urlDetails:=Url{}

	db := database.Connect()
	defer db.Close()

	rows,err:=db.Query("SELECT * FROM urls WHERE shortened = $1",url)
	if err != nil{
		fmt.Println(err)
		return c.Status(fiber.StatusNoContent).JSON(&fiber.Map{
			"success":false,
			"data":"Invalid url",
		})
	}
	for rows.Next(){
		err := rows.Scan(&urlDetails.Id,&urlDetails.Original,&urlDetails.Shortened,&urlDetails.User_id,&urlDetails.RateRemaining,urlDetails.Expiry,&urlDetails.RateLimitReset,&urlDetails.IsActive)
		if err != nil{
			fmt.Println(err)
			return c.Status(400).JSON(&fiber.Map{
				"success":false,
				"data":"Error scaning data",
			})
		}
	}

	if !urlDetails.IsActive{
		return c.Status(300).JSON(&fiber.Map{
			"success":false,
			"data":"Url is not active",
		})
	}

	if urlDetails.RateRemaining == 0 {
		return c.Status(300).JSON(&fiber.Map{
			"success":false,
			"data":"Url rate limit reached",
		})
	}


	return c.Redirect(urlDetails.Original,302)
}