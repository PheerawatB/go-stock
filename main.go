package main

import (
	"fmt"
	"gostock/handlers"
	"gostock/repositories"
	"gostock/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	repo := repositories.NewOrderRepository()
	service := services.NewOrderService(repo)

	app := fiber.New()

	app.Post("/orders", func(c *fiber.Ctx) error { return handlers.CreateOrderHandler(service)(c) })
	app.Get("/orders/:id", func(c *fiber.Ctx) error { return handlers.GetOrderHandler(repo)(c) })
	app.Get("/orders", func(c *fiber.Ctx) error { return handlers.GetAllOrdersHandler(repo)(c) })
	app.Delete("/orders/:id", func(c *fiber.Ctx) error { return handlers.DeleteOrderHandler(repo)(c) })

	port := 8080
	fmt.Printf("Server is running on port %d\n", port)
	app.Listen(fmt.Sprintf(":%d", port))
}
