package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hritesh04/url-shortner/database"
	"github.com/hritesh04/url-shortner/middleware"
	"github.com/hritesh04/url-shortner/prometheus"
	"github.com/hritesh04/url-shortner/routes"
	_ "github.com/lib/pq"
	"github.com/lpernett/godotenv"
)


func setupRoutes(app *fiber.App){
	app.Get("/stats", prometheus.GetStats)
	app.Post("/signin", routes.SignIn)
	app.Post("/signup", routes.SignUp)
	app.Get("/:url", routes.Resolve)

	app.Get("/user/details",middleware.UserAuth, routes.GetUserDetails)
	app.Post("generateQr",middleware.UserAuth,routes.GenerateQr)
	app.Post("/shorten",middleware.UserAuth, routes.Shorten)
}


func main() {
	
	if err := godotenv.Load(); err!=nil{
		log.Fatal("Failed to Load ENV variables")
	}
	
	if err := database.Init(); err!=nil{
		log.Fatal(err)
	}

	
	app:= fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowCredentials: true,
	}))
	
	app.Use(logger.New())
	
	prometheus.InitMetrics()
	app.Get("/metrics", prometheus.Metrics)
	
	database.InitCache()
	
	setupRoutes(app)
	
	app.Listen(os.Getenv("PORT"))
}