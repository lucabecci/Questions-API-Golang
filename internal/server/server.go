package server

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	v1 "github.com/lucabecci/questions-golang-API/internal/server/v1"
)

type Server struct {
	server *fiber.App
}

type Versions struct {
	V1 string
}

func GetInstance() (*Server, error) {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Question",
		ReadTimeout:   10 * time.Second,
		WriteTimeout:  10 * time.Second,
	})

	//routes
	v1 := v1.New()
	app.Mount("/v1", v1)
	//instance
	server := Server{server: app}
	//return srv
	return &server, nil
}

func (s *Server) Start(port string) error {
	err := s.server.Listen(":" + port)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Close() error {
	s.Close()
	fmt.Println("Server Stoped")
	return nil
}

func Index(c *fiber.Ctx) {
	var versions = Versions{
		V1: "http://localhost:4000/v1",
	}
	result, err := json.Marshal(versions)
	if err != nil {
		log.Panic(err.Error())
		return
	}
	c.Send(result)
	return
}
