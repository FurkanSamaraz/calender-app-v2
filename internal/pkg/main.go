// @title Your Project API
// @description This is the API for your project
// @version 1
// @host localhost:8080
// @BasePath /api/v1

package main

import (
	"log"
	apps "main/external-api"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetCollection() *gorm.DB {
	dsn := "host=localhost user=postgres password=172754 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
func main() {
	app := fiber.New()

	apps.Setup(app, GetCollection())

	app.Listen(":3000")
}
