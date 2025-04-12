package main

import (
	"github.com/NewChakrit/golang_hexagonal/adapters"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/NewChakrit/golang_hexagonal/core"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// initialize the database connection
	// filed based
	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := core.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	app.Post("/order", orderHandler.CreateOrder)
	// migrate the schema
	db.AutoMigrate(&core.Order{})

	app.Listen(":8080")
}
