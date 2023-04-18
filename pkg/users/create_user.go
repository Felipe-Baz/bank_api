package users

import (
	"bank_api/pkg/commons/models"

	"github.com/gofiber/fiber/v2"
)

type CreateUserRequestBody struct {
    Document string    `json:"document"`
    Email string       `json:"email"`
    Name  string       `json:"name"`
    Birthdate string   `json:"birthdate"`
}

// @Summary Create User
// @Description function to create a new user in database
// @Tags users
// @Produce json
// @Success 200  {object} models.User
// @Router /user [post]
func (h handler) CreateUser(context *fiber.Ctx) error {
	var user models.User

	err_factory := userFactory(context, &user)
	if err_factory != nil {
		return err_factory
	}

	err_create := h.userCreateInDB(&user)

	if err_create != nil {
		return err_create
	}

	return context.Status(fiber.StatusCreated).JSON(&user)
}

func userFactory(context *fiber.Ctx, user *models.User) *fiber.Error {
	body := CreateUserRequestBody{}

	// Parse body
	err := context.BodyParser(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body.userConvert(user)
	return nil
}

func (h handler) userCreateInDB(user *models.User) *fiber.Error {
	result := h.DB.Create(&user)

	if result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}

	return nil
}

func (body CreateUserRequestBody) userConvert(user *models.User) {
	user.Document = body.Document
	user.Email = body.Email
	user.Name = body.Name
	user.Birthdate = body.Birthdate
}