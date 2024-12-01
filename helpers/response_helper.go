package helpers

import "github.com/gofiber/fiber/v2"

func Success(c *fiber.Ctx, message string) error {
	return c.Status(200).JSON(fiber.Map{
		"message": message,
	})
}

func Error(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(fiber.Map{
		"message": message,
	})
}
