package users

import (
	"bank_api/pkg/commons/models"

	"github.com/gofiber/fiber/v2"
)

type UpdateUserRequestBody struct {
    Document string    `json:"document"`
    Email string       `json:"email"`
    Name  string       `json:"name"`
    Birthdate string   `json:"birthdate"`
}

func (h handler) UpdateUser(context *fiber.Ctx) error {
    id := context.Params("id")
    body := UpdateUserRequestBody{}

    // getting request's body
    if err := context.BodyParser(&body); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, err.Error())
    }

    var user models.User

    if result := h.DB.First(&user, id); result.Error != nil {
        return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
    }

    body.userConvert(&user)

    // save product
    h.DB.Save(&user)

    return context.Status(fiber.StatusOK).JSON(&user)
}

func (body UpdateUserRequestBody) userConvert(user *models.User) {
	user.Document = body.Document
	user.Email = body.Email
	user.Name = body.Name
	user.Birthdate = body.Birthdate
}