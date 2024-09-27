package rest

import (
	"context"
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hritesh04/url-shortner/internal/dto"
)

type Cache interface {
	Get(string) (string, error)
	Set(context.Context, string, string, time.Duration) error
}

type Auth interface {
	GetUserData(string) (*dto.Users, error)
	Authorize(*fiber.Ctx) error
	HashPassword(string) string
	ComparePassword(string, string) bool
	GenerateToken(int32, string) (string, error)
}

type Monitor interface {
	Metrics(*fiber.Ctx) error
	GetStats(string, string, string) ([]byte, error)
	Increment(string, string)
}

type RestHandler struct {
	App     *fiber.App
	DB      *sql.DB
	Cache   Cache
	Auth    Auth
	Monitor Monitor
}
