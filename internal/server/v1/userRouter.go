package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucabecci/questions-golang-API/internal/database/data"
	"github.com/lucabecci/questions-golang-API/internal/helpers"
	"github.com/lucabecci/questions-golang-API/internal/server/middlewares"
)

type UserRouter struct {
	Repository data.UserRepository
}

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ur *UserRouter) Register(c *fiber.Ctx) error {
	var body UserInput
	err := c.BodyParser(&body)
	if err != nil {
		return fiber.NewError(400, "Error in the camps")
	}
	exists := ur.Repository.EmailInUse(body.Email)
	if exists == true {
		return fiber.NewError(400, "Email already in use")
	}
	usr, succes := ur.Repository.Create(body.Email, body.Password)
	if succes == false {
		return fiber.NewError(400, "Error to create your account")
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": struct {
			ID       int    `json:"id"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}{
			ID:       int(usr.ID),
			Email:    usr.Email,
			Password: usr.Password,
		},
	})
	return nil
}

func (ur *UserRouter) Login(c *fiber.Ctx) error {
	var body UserInput
	err := c.BodyParser(&body)
	if err != nil {
		return fiber.NewError(400, "Error in the camps")
	}
	usr, err := ur.Repository.UserExists(body.Email)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}
	result := usr.PasswordMatch(body.Password)
	if result == false {
		return fiber.NewError(400, "Password invalid")
	}
	token, err := helpers.JwtGenerator(usr.ID)
	if err != nil {
		return fiber.NewError(500, "Internal Server Error")
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
		"user": struct {
			ID    int    `json:"id"`
			Email string `json:"email"`
		}{
			ID:    int(usr.ID),
			Email: usr.Email,
		},
	})
	return nil
}

func (ur *UserRouter) Account(c *fiber.Ctx) error {
	uID, err := helpers.ConvertMetaData(c)
	if err != nil {
		return fiber.NewError(500, "Internal server error")
	}
	usr, success := ur.Repository.GetOne(uID)
	if success == false {
		return fiber.NewError(500, "Internal server error")
	}
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hello " + usr.Email,
	})
	return nil
}

func (ur *UserRouter) Service() *fiber.App {
	service := fiber.New()

	service.Post("/register", ur.Register)
	service.Post("/login", ur.Login)
	service.Get("/account", middlewares.Protected(), ur.Account)
	return service
}
