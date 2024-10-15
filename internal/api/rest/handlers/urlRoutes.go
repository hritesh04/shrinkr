package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hritesh04/url-shortner/internal/api/rest"
	"github.com/hritesh04/url-shortner/internal/dto"
	"github.com/hritesh04/url-shortner/internal/repository"
	"github.com/hritesh04/url-shortner/internal/service"
)

type UrlHandler struct {
	svc service.UrlService
}

func SetupUrlRoutes(rh *rest.RestHandler) {

	app := rh.App

	svc := service.UrlService{
		Repo:    repository.NewUrlRepository(rh.DB, rh.Cache),
		Auth:    rh.Auth,
		Monitor: rh.Monitor,
	}

	handler := UrlHandler{
		svc: svc,
	}

	app.Get("/:url", handler.Resolve)

	pvtGroup := app.Group("/", rh.Auth.Authorize)
	pvtGroup.Post("/shorten", handler.Shorten)

}

func (u *UrlHandler) Resolve(ctx *fiber.Ctx) error {
	postfix := ctx.Params("url")
	url, err := u.svc.Resolve(postfix)
	if err != nil {
		return ctx.Status(404).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	return ctx.Status(301).Redirect(url)
}

func (u *UrlHandler) Shorten(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(dto.Claim)
	body := dto.Request{}
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   "err",
		})
	}

	url, err := u.svc.ShortenUrl(&body, &user)
	if err != nil {
		return ctx.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"success": true,
		"data":    url,
	})

}
