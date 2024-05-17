package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	jtoken "github.com/golang-jwt/jwt/v4"
)

type Claim struct {
	jtoken.RegisteredClaims
	Id int32
}

func UserAuth(c *fiber.Ctx) error{

	var userClaim Claim

	token:=c.Cookies("urlshortTkn")
	if token == ""{
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"success":true,
			"data":"Auth token missing",
		})
	}

	tokenByte,err := jtoken.ParseWithClaims(token,&userClaim, func(t *jtoken.Token) (interface{}, error) {
		if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")),nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"success": false,
			"data":    "Invalid token",
		})
	}

	if !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"success": false,
			"data":    "Invalid token",
		})
	}

	userId:= userClaim.Id

	c.Locals("userId", userId)

	return c.Next()
}