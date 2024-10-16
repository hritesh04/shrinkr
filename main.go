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

func SetupEnv() (api.AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return api.AppConfig{}, errors.New("failed to load")
	}

	port := os.Getenv("PORT")

	if len(port) < 1 {
		log.Println("port not found, using default port 3000")
		port = ":3000"
	}

	Dbn := os.Getenv("DB_CONNSTR")
	if len(Dbn) < 1 {
		return api.AppConfig{}, errors.New("database url not found")
	}

	PROM_URL := os.Getenv("PROM_URL")
	if len(PROM_URL) < 1 {
		return api.AppConfig{}, errors.New("prometheus url not found")
	}

	secret := os.Getenv("SECRET")
	if len(secret) < 1 {
		return api.AppConfig{}, errors.New("app secret not found")
	}

	return api.AppConfig{DB_Str: Dbn, Port: port, Secret: secret,
		PROM_URL: PROM_URL,
		Site_url: os.Getenv("SITE_URL"),
		Sub_free: os.Getenv("SUB_FREE"),
		Sub_pre:  os.Getenv("SUB_PRE")}, nil
}
