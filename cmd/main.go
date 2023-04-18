package main

import (
    "log"

    "bank_api/pkg/commons/config"
    "bank_api/pkg/commons/db"
    "bank_api/pkg/users"

    "github.com/gofiber/fiber/v2"
)

func main() {
    configs, err := config.LoadConfig()

    if err != nil {
        log.Fatalln("Failed to load configs", err)
    }

    database := db.Init(&configs)
    app := fiber.New()

    users.RegisterRoutes(app, database)

    app.Listen(configs.Port)
}