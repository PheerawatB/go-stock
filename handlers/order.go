package handlers

import (
	"gostock/repositories"
	"gostock/services"

	"github.com/gofiber/fiber/v2"
)

func CreateOrderHandler(service *services.OrderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var order repositories.Order
		if err := c.BodyParser(&order); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		service.ProcessOrder(order)
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Order is being processed"})
	}
}

func GetOrderHandler(repo *repositories.OrderRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id") // Get the ID from URL params
		order, exists := repo.FindByID(id)
		if !exists {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Order not found"})
		}
		return c.JSON(order)
	}
}

func GetAllOrdersHandler(repo *repositories.OrderRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		orders := repo.FindAll()
		return c.JSON(orders)
	}
}

func DeleteOrderHandler(repo *repositories.OrderRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		repo.DeleteByID(id)
		return c.SendStatus(fiber.StatusNoContent)
	}
}
