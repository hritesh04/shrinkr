package api

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hritesh04/url-shortner/internal/api/rest"
	"github.com/hritesh04/url-shortner/internal/api/rest/handlers"
	"github.com/hritesh04/url-shortner/internal/database"
	"github.com/hritesh04/url-shortner/internal/helper"
	monitor "github.com/hritesh04/url-shortner/pkg/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)



func SetupServer(cfg interface{}) {
	app := fiber.New()

	app.Use(logger.New())

	db,err := database.Init(); if err != nil{
		log.Fatalf("database connection error %v\n", err)
	}

	cache := database.InitCache()
	auth := helper.SetupAuth(os.Getenv("SECRET"))
	
	topics := map[string]*prometheus.CounterVec{
		"UrlVisitCount":prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "url_visit_count",
				Help: "Total number of times URL is visited",
			},
			[]string{"url"},
		),
	}
	
	monitor := monitor.InitMetrics(topics)


	rh := &rest.RestHandler{
		App: app,
		DB: db,
		Cache: cache,
		Auth: auth,
		Monitor : monitor,
	}

	setupRoutes(rh)

	app.Listen(os.Getenv("PORT"))
}

func setupRoutes(rh *rest.RestHandler){
	handlers.SetupMetricsRoute(rh)
	handlers.SetupUrlRoutes(rh)
}