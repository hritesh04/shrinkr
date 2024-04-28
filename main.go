package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/hritesh04/url-shortner/database"
	"github.com/hritesh04/url-shortner/routes"
	_ "github.com/lib/pq"
	"github.com/lpernett/godotenv"
)

func setupRoutes(app *fiber.App){
	app.Post("/signin", routes.SignIn)
	app.Post("/signup", routes.SignUp)
	app.Get("/user/:userId", routes.GetUserDetails)
	app.Get("/:url", routes.Resolve)
}


func main() {

	if err := godotenv.Load(); err!=nil{
		log.Fatal("Failed to Load ENV variables")
	}
	
	if err := database.Init(); err!=nil{
			log.Fatal(err)
	}

	app:= fiber.New()

	setupRoutes(app)
	
	app.Listen(os.Getenv("PORT"))
}