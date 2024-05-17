package helper

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/hritesh04/url-shortner/internal/dto"
)

type Auth struct {
	Secret string
}

func SetupAuth(secret string) *Auth {
	return &Auth{
		Secret: secret,
	}
}

func (a *Auth)GetUserData(token string)(*dto.Users,error){
	
	var user dto.Claim

	tokenByte, err := jtoken.ParseWithClaims(token, &user, func(t *jtoken.Token) (interface{}, error) {
		if _, ok := t.Method.(*jtoken.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil{
		return &dto.Users{},fmt.Errorf("error parsing the token")
	}

	if !tokenByte.Valid {
		return &dto.Users{},fmt.Errorf("invalid token")
	}

	userData := &dto.Users{
		Id:					user.Id,
		SubscriptionType: user.SubscriptionType,
	}

	return userData,nil

}

func (a *Auth) Authorize(c *fiber.Ctx)error {

	var userClaim dto.Claim

	token := c.Cookies("urlshortTkn")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"success": true,
			"data":    "Auth token missing",
		})
	}

	tokenByte, err := jtoken.ParseWithClaims(token, &userClaim, func(t *jtoken.Token) (interface{}, error) {
		if _, ok := t.Method.(*jtoken.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
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
	c.Locals("user",userClaim)
	return c.Next()
}