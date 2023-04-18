package users

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

type handler struct {
    DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
    h := &handler{
        DB: db,
    }

    routes := app.Group("/user")
    routes.Post("/", h.CreateUser)
    routes.Get("/", h.GetUsers)
    // routes.Get("/:id", h.GetUserById)
    // routes.Put("/:id", h.UpdateUser)
    // routes.Delete("/:id", h.DeleteUser)
}