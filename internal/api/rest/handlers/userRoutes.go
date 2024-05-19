package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hritesh04/url-shortner/internal/api/rest"
	"github.com/hritesh04/url-shortner/internal/dto"
	"github.com/hritesh04/url-shortner/internal/repository"
	"github.com/hritesh04/url-shortner/internal/service"
)

type UserHandler struct {
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler){ 
	app:= rh.App

	svc:= service.UserService{
		Repo: repository.NewUserRepository(rh.DB),
		Auth: rh.Auth,
	}

	handler := UserHandler{
		svc:svc,
	}

	app.Post("/signup",handler.Signup)
	app.Post("/signin",handler.Signin)

	pvtGroup := app.Group("/user", handler.svc.Auth.Authorize)
	pvtGroup.Get("/",handler.GetUserDetails)


}

func (uh *UserHandler)Signin(ctx *fiber.Ctx)error{
	reqForm := dto.SignInRequest{}
	if err := ctx.BodyParser(&reqForm); err != nil {
		return ctx.Status(500).JSON(&fiber.Map{
			"success":true,
			"error":err,
		})
	}
	token,err := uh.svc.Login(&reqForm)

	if err != nil {
		return ctx.Status(500).JSON(&fiber.Map{
			"success":false,
			"error":err,
		})
	}

	cookie := new(fiber.Cookie)
	cookie.Name="shrinkr"
	cookie.Value=token
	cookie.Expires=time.Now().Add(24 * time.Hour)
	ctx.Cookie(cookie)

	return ctx.Status(200).JSON(&fiber.Map{
		"success":true,
		"token":token,
	})
}

func (uh *UserHandler)Signup(ctx *fiber.Ctx)error{
	reqForm := dto.SignUpRequest{}
	if err := ctx.BodyParser(&reqForm); err != nil {
		return ctx.Status(500).JSON(&fiber.Map{
			"success":false,
			"error":err,
		})
	}
	token,err := uh.svc.CreateUser(&reqForm)
	if err != nil {
		return ctx.Status(500).JSON(&fiber.Map{
			"succes":false,
			"error":err,
		})
	}
	cookie := new(fiber.Cookie)
	cookie.Name="shrinkr"
	cookie.Value=token
	cookie.Expires=time.Now().Add(24 * time.Hour)
	ctx.Cookie(cookie)

	return ctx.Status(200).JSON(&fiber.Map{
		"success":true,
		"token":token,
	})
}

func (uh *UserHandler)GetUserDetails(ctx *fiber.Ctx)error{
	
	token := ctx.Locals("user").(dto.Claim)
	
	user,err := uh.svc.GetUserDetails(token.Id)

	if err != nil {
		return ctx.Status(500).JSON(&fiber.Map{
			"success":false,
			"error":err,
		})
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"success":true,
		"data":user,
	})
}