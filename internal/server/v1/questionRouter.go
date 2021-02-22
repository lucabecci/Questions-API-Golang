package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucabecci/questions-golang-API/internal/database/data"
	"github.com/lucabecci/questions-golang-API/internal/helpers"
	"github.com/lucabecci/questions-golang-API/internal/server/middlewares"
)

type QuestionRouter struct {
	Repository data.QuestionRepository
}

type QuestionInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ID          uint
}

func (qr *QuestionRouter) CreateQuestion(c *fiber.Ctx) error {
	uID, err := helpers.ConvertMetaData(c)
	if err != nil {
		return fiber.NewError(500, "Internal server error")
	}
	var body QuestionInput
	err = c.BodyParser(&body)
	if err != nil {
		return fiber.NewError(400, "Error in the camps")
	}
	result, success := qr.Repository.Create(data.Question{
		Title:       body.Title,
		Description: body.Description,
		UserID:      uID,
	})
	if success == false {
		return fiber.NewError(500, "Error to create your question")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"question": result,
	})
}

func (qr *QuestionRouter) GetQuestions(c *fiber.Ctx) error {
	uID, err := helpers.ConvertMetaData(c)
	if err != nil {
		return fiber.NewError(500, "Internal server error")
	}
	results, success := qr.Repository.GetByUser(uID)
	if success == false {
		return fiber.NewError(400, "You dont have questions")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"questions": results,
	})
}

func (qr *QuestionRouter) Service() *fiber.App {
	service := fiber.New()
	service.Get("/", middlewares.Protected(), qr.GetQuestions)
	service.Post("/", middlewares.Protected(), qr.CreateQuestion)
	return service
}
