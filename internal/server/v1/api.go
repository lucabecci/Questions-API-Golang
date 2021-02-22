package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucabecci/questions-golang-API/internal/database"
	"github.com/lucabecci/questions-golang-API/internal/database/data"
)

func New() *fiber.App {
	router := fiber.New()
	db := database.Database()
	ur := UserRouter{Repository: data.UserRepository{Database: db}}
	qr := QuestionRouter{Repository: data.QuestionRepository{Database: db}}
	userRouter := ur.Service()
	questionRouter := qr.Service()

	router.Mount("/user", userRouter)
	router.Mount("/question", questionRouter)
	return router
}
