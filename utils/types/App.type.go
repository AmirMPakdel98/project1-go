package types

import (
	"c-vod/utils/storage"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type App struct {
	Config  *AppConfig
	Server  *fiber.App
	Router  fiber.Router
	DB      *gorm.DB
	Storage *storage.Storage
}
