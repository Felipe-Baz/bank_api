package users

import (
	"bank_api/pkg/commons/entity"

	"github.com/gofiber/fiber/v2"
)

func (h handler) DeleteUser(context *fiber.Ctx) error {
	id := context.Params("id")

	var user entity.User

	if result := h.DB.First(&user, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	// delete product from db
	h.DB.Delete(&user)

	return context.SendStatus(fiber.StatusOK)
}
