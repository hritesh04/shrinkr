package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hritesh04/url-shortner/internal/api/rest"
	"github.com/hritesh04/url-shortner/internal/api/rest/handlers"
	"github.com/hritesh04/url-shortner/internal/database"
	"github.com/hritesh04/url-shortner/internal/helper"
	monitor "github.com/hritesh04/url-shortner/pkg/prometheus"
)

type AppConfig struct {
	DB_Str   string
	Port     string
	Secret   string
	Site_url string
	Sub_free string
	Sub_pre  string
}

func SetupServer(cfg AppConfig) {
	app := fiber.New()

	app.Use(logger.New())

	db, err := database.Init(cfg.DB_Str)
	if err != nil {
		log.Fatalf("database connection error %v\n", err)
	}

	cache := database.InitCache()
	auth := helper.SetupAuth(cfg.Secret)

	monitor := monitor.NewMonitorService()

	rh := &rest.RestHandler{
		App:     app,
		DB:      db,
		Cache:   cache,
		Auth:    auth,
		Monitor: monitor,
	}

	setupRoutes(rh)

	app.Listen(cfg.Port)
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupMetricsRoute(rh)
	handlers.SetupUrlRoutes(rh)
	handlers.SetupUserRoutes(rh)
}
