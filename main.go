package main

import (
	"errors"
	"log"
	"os"

	"github.com/hritesh04/url-shortner/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	cfg, err := SetupEnv()
	if err != nil {
		log.Fatalf("config file is not loaded properly %v\n", err)
	}
	api.SetupServer(cfg)
}


type AppConfig struct {
	DB_Str		string
	Port       string
	Secret     string
	Site_url   string
	Sub_free   string
	Sub_pre    string
}

func SetupEnv()(AppConfig,error) {
	err := godotenv.Load()
	if err != nil {
		return AppConfig{},errors.New("failed to load")
	}

	port := os.Getenv("PORT")

	if len(port) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	Dbn := os.Getenv("DB_CONNSTR")
	if len(Dbn) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	secret := os.Getenv("SECRET")
	if len(secret) < 1 {
		return AppConfig{}, errors.New("app secret not found")
	}

	return AppConfig{DB_Str: Dbn,Port:port,Secret: secret,
		Site_url: 	os.Getenv("SITE_URL"),
		Sub_free: 	os.Getenv("SUB_FREE"),
		Sub_pre: 	os.Getenv("SUB_PRE"),},nil
}