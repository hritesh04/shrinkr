package routes

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"

	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/hritesh04/url-shortner/database"
	"github.com/hritesh04/url-shortner/helper"
	"github.com/hritesh04/url-shortner/models"
)


func GetUserDetails(c *fiber.Ctx)error {
	userid:=c.Locals("userId")
	user := models.Users{}
	
	db:= database.Connect()
	defer db.Close()
	
	rows, err := db.Query("SELECT * FROM USERS WHERE id = $1;", userid)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success":false,
			"data":"",
		})
	}	
	defer rows.Close()
	
	for rows.Next(){
		err := rows.Scan(&user.Id,&user.Name,&user.Email,&user.Password,&user.SubscriptionType);
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
	body := models.Users{}
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
	
	var id int32
	err := db.QueryRow("INSERT INTO USERS (name,email,password) VALUES ($1,$2,$3) RETURNING id;", body.Name, body.Email, body.Password).Scan(&id)

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success":false,
			"data":err,
		})
	}

	// create a token and put in cookie
	data := jtoken.NewWithClaims(jtoken.SigningMethodHS256,models.Claim{
		RegisteredClaims: jtoken.RegisteredClaims{},
		Id:id,
	})
	
	token,err := data.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return	c.Status(500).JSON(&fiber.Map{
			"success":false,
			"data":data,
		})
	}

	cookie := new(fiber.Cookie)
	cookie.Name="urlshortTkn"
	cookie.Value=token
	cookie.Expires=time.Now().Add(24* time.Hour)

	c.Cookie(cookie)

	return c.Status(200).JSON(&fiber.Map{
		"success":true,
		"data":token,
	})
} 

func SignIn(c *fiber.Ctx)error {
	body := models.SignUpRequest{}

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
	user := models.Users{}
	for rows.Next(){
		err := rows.Scan(&user.Id,&user.Name,&user.Email,&user.Password,&user.SubscriptionType)
		if err != nil{
			return c.Status(500).JSON(&fiber.Map{
				"success":false,
				"data":err,
			})
		}
	}

	correctPassword := helper.ComparePassword(user.Password,body.Password)

	if correctPassword {

		// create token and put in cookie

		data := jtoken.NewWithClaims(jtoken.SigningMethodHS256,models.Claim{
			RegisteredClaims: jtoken.RegisteredClaims{},
			Id:user.Id,
		})

		token,err := data.SignedString([]byte(os.Getenv("SECRET")))
		if err != nil {
			return	c.Status(500).JSON(&fiber.Map{
				"success":false,
				"data":"Failed to generate token",
			})
		}

		cookie := new(fiber.Cookie)
		cookie.Name="urlshortTkn"
		cookie.Value=token
		cookie.Expires=time.Now().Add(24 * time.Hour)

		c.Cookie(cookie)

		return c.Status(200).JSON(&fiber.Map{
			"success":true,
			"data":token,
		})
	}else{
		return c.Status(400).JSON(&fiber.Map{
			"success":false,
			"data":"Incorrect Password",
		})
	}
}