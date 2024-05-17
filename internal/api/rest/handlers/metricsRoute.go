package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hritesh04/url-shortner/internal/api/rest"
)

type MonitorHandler struct{
	svc	rest.Monitor
}

func SetupMetricsRoute(rh *rest.RestHandler) {
	app := rh.App

	handler := MonitorHandler{
		svc:rh.Monitor,
	}

	app.Get("/metrics", handler.Metrics)
	app.Get("/stats", handler.GetStats)
}

func (m *MonitorHandler)Metrics(ctx *fiber.Ctx)error{
	return m.svc.Metrics(ctx)
}

func (m *MonitorHandler)GetStats(ctx *fiber.Ctx)error{
	url := ctx.Query("url")
	step := ctx.Query("step")
	limit := ctx.Query("limit")
	data,err := m.svc.GetStats(url,step,limit)
	if err != nil {
		return ctx.Status(404).JSON(&fiber.Map{
			"success":false,
			"error":err,
		})
	}
	return ctx.Status(200).JSON(&fiber.Map{
		"success":true,
		"data":data,
	})
}