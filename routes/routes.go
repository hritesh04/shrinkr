package routes

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hritesh04/url-shortner/database"
	"github.com/hritesh04/url-shortner/helper"
)

type Users struct{
	Id int32
	Name string
	Email string
	Password string
	Urls []Url
}

type Url struct{
	Id int32 
	Original string
	Shortened string
	User_id int32
	RateRemaining int32
	Expiry time.Time
	RateLimitReset time.Time
	IsActive bool
}


func GetUserDetails(c *fiber.Ctx)error {
	userid:=c.Params("userId")

	user := Users{}
	
	db:= database.Connect()
	defer db.Close()
	
	rows, err := db.Query(`
		SELECT * FROM USERS WHERE id = $1;`, userid)
	if err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success":false,
			"data":"",
		})
	}	
	defer rows.Close()
	
	for rows.Next(){
		err := rows.Scan(&user.Id,&user.Name,&user.Email,&user.Password);
		if err != nil {
			fmt.Println(err)
			return c.Status(400).JSON(&fiber.Map{
				"success":false,
				"data":"",
			})
		}
	}

	return c.Status(200).JSON(&fiber.Map{
		"success":true,
		"data":user,
	})
}

func SignUp(c *fiber.Ctx)error {
	body := Users{}
	if err := c.BodyParser(&body); err != nil{
		return c.Status(500).JSON(&fiber.Map{
			"success":false,
			"data":"Error while parsing body",
		})
	}

	hash := helper.EncodePassword(body.Password)

	if hash == ""{
		return c.Status(500).JSON(&fiber.Map{
			"success":false,
			"data":"Error while hashing password",
		})
	}
	
	body.Password = hash

	db := database.Connect()
	defer db.Close()

	result ,err := db.Exec("INSERT (name,email,password) INTO USERS VALUES ($1,$2,$3)",body.Name,body.Email,body.Password)

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success":false,
			"data":"Error Creating USer",
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success":true,
		"data":result,
	})
} 

func SignIn(c *fiber.Ctx)error {
	body := Users{}

	if err := c.BodyParser(&body);err != nil{
		return c.Status(400).JSON(&fiber.Map{
			"success":false,
			"data":"Error parsing json",
		})
	}

	db := database.Connect()
	defer db.Close()

	rows,err := db.Query("SELECT * FROM USERS WHERE email = $1",body.Email)

	if err != nil{
		return c.Status(500).JSON(&fiber.Map{
			"success":false,
			"data":"User not found",
		})
	}
	user := Users{}
	for rows.Next(){
		err := rows.Scan(&user.Id,&user.Name,&user.Email,&user.Password)
		if err != nil{
			return c.Status(500).JSON(&fiber.Map{
				"success":false,
				"data":"User mapping failed",
			})
		}
	}

	correctPassword := helper.ComparePassword(user.Password,body.Password)

	if correctPassword {
		return c.Status(200).JSON(&fiber.Map{
			"success":true,
			"data":user,
		})
	}else{
		return c.Status(400).JSON(&fiber.Map{
			"success":false,
			"data":"Incorrect Password",
		})
	}

}