package users

import (
	"bank_api/pkg/commons/entity"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get Users
// @Description function to list all users in database
// @Tags users
// @Produce json
// @Success 200  {object}  []entity.User
// @Router /user [get]
func (h handler) GetUsers(context *fiber.Ctx) error {
	var users []entity.User

	err := h.getUsersInDB(&users)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(&users)
}

func (h handler) getUsersInDB(users *[]entity.User) *fiber.Error {
	result := h.DB.Find(users)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	return nil
}
