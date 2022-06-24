package main

import (
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

// TODO: do more things, yknow?
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		cmd, err := exec.Command("/bin/sh", "./update.sh").Output()

		if err != nil {
			return c.Status(500).SendString("Internal Server Error")
		}

		return c.JSON(fiber.Map{
			"output": string(cmd),
		})
	})

	app.Listen(":5000")
}
