package external_api

import (
	retail "main/retail"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app fiber.Router, db *gorm.DB) {
	retail.Run(app, db)

}
