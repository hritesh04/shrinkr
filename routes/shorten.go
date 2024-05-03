package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/hritesh04/url-shortner/database"
	"github.com/hritesh04/url-shortner/helper"
	"github.com/hritesh04/url-shortner/models"
)



func Shorten(c *fiber.Ctx)error{
	userId := c.Locals("userId")
	user:=models.Users{}
	urlData := models.Url{}
	body:=models.Request{}
	if err := c.BodyParser(&body); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":false,
			"data":"can not parse JSON",
		})
	}

	fmt.Printf("%+v\n", body)

	db := database.Connect()
	defer db.Close()

	// check user plan using userId

	// ** add subscription coloum and check subscription

	rows,err := db.Query("SELECT * FROM USERS WHERE id = $1",userId)

	if err != nil{
		return c.Status(500).JSON(&fiber.Map{
			"success":false,
			"data":"User not found",
		})
	}

	for rows.Next(){
		err := rows.Scan(&user.Id,&user.Name,&user.Email,&user.Password,&user.SubscriptionType)
		if err != nil{
			return c.Status(500).JSON(&fiber.Map{
				"success":false,
				"data":"User mapping failed",
			})
		}
	}

	
	
	// set expiry according to the plan
	
	
	
	// set Rate remaining according to the plan
	urlData.Original=body.Url
	urlData.User_id=user.Id
	
	if body.CustomUrl == "" {
		urlData.Shortened = uuid.New().String()[:6]
		}else{
			urlData.Shortened = body.CustomUrl
		}
		
		// add to db and send response
		
	fmt.Printf("%+v\n", urlData)
	data,err := helper.AddUrl(&urlData)
	if err != nil{
		return c.Status(500).JSON(&fiber.Map{
			"success":false,
			"data":err,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success":true,
		"data":data,
	})
}