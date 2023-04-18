package users

import (
	"bank_api/pkg/commons/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetUserById(context *fiber.Ctx) error {
	id := context.Params("id")

    var user models.User

	h.getUserById(context, &user, id)

    return context.Status(fiber.StatusOK).JSON(&user)
}

func (h handler) getUserById(context *fiber.Ctx, user *models.User, id string) error {
	
	result := h.DB.First(&user, id)
    if result.Error != nil {
        return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
    }
	return nil

}