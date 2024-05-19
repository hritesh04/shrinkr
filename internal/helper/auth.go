package helper

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/hritesh04/url-shortner/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Secret string
}

func SetupAuth(secret string) *Auth {
	return &Auth{
		Secret: secret,
	}
}

func (a *Auth)HashPassword(password string)string{
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return ""
	}	
	return string(hashedPassword)
}

func (a *Auth)ComparePassword(password string,hash string)bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	if err == nil{
		return true
	}
	return false
}

func (a *Auth)GenerateToken(userId int32,subscription string)(string,error){
	data := jwt.NewWithClaims(jtoken.SigningMethodHS256,dto.Claim{
		RegisteredClaims: jtoken.RegisteredClaims{},
		Id: userId,
		SubscriptionType: subscription,
	})

	token,err := data.SignedString([]byte(a.Secret))
	if err != nil {
		return "",err
	}
	return token,nil
}

func (a *Auth)GetUserData(token string)(*dto.Users,error){
	
	var user dto.Claim

	tokenByte, err := jtoken.ParseWithClaims(token, &user, func(t *jtoken.Token) (interface{}, error) {
		if _, ok := t.Method.(*jtoken.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(a.Secret), nil
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

	token := c.Cookies("shrinkr")
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
		return []byte(a.Secret), nil
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