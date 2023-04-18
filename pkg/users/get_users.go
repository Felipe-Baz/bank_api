package users

import (
    "bank_api/pkg/commons/models"

    "github.com/gofiber/fiber/v2"
)

func (h handler) GetUsers(c *fiber.Ctx) error {
    var users []models.User

    err := h.getUsersInDB(&users)
	if err != nil {
		return err
	}

    return c.Status(fiber.StatusOK).JSON(&users)
}

func (h handler) getUsersInDB(users *[]models.User) *fiber.Error {
	result := h.DB.Find(users)
	if result.Error != nil {
        return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
    }
	return nil
}