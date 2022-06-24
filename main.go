package main

import (
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	code := getCode()

	app.Get("/", func(c *fiber.Ctx) error {
		if c.GetReqHeaders()["Authorization"] != code {
			return c.Status(401).SendString("Unauthorized")
		}

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

func getCode() string {
	content, err := os.ReadFile("code.txt")
	if err != nil {
		return "password"
	}
	return string(content)
}
