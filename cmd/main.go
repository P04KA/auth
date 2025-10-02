package main

import (
	"context"

	"github.com/P04KA/auth/internal/app"
	"github.com/P04KA/auth/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	key = "thedoghousecasinogambling"
)

var conn *pgxpool.Pool
var ctx = context.Background()

func isTokenValid(c *fiber.Ctx) error {
	header := c.Get("Authorization")

	if header[:7] != "Bearer " {
		return c.JSON(fiber.Map{"authorized": false})
	}

	token := header[7:]

	if token == key {
		var user models.User
		err := conn.QueryRow(ctx, "SELECT * FROM auth LIMIT 100").Scan(&user.Email, &user.Name)

		if errors.Is(err, pgx.ErrNoRows) {
			return c.JSON(fiber.Map{"authorized": false, "error": "user not found"})
		}
		if err != nil {
			errors.Wrap(err, "db error")
			return c.JSON(fiber.Map{"authorized": false, "error": "database error"})
		}

		return c.JSON(fiber.Map{
			"authorized": true,
			"email":      user.Email,
			"name":       user.Name,
		})
	}
	return c.JSON(fiber.Map{"authorized": false})
}

func main() {
	if err := app.Run(); err != nil {
		logrus.WithError(err).Fatal("Failed to start application")
		return
	}
	defer conn.Close()

	app := fiber.New()
	app.Get("/profile", isTokenValid)
	app.Listen(":8080")
}
