package main

import (
	"log"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"
	"github.com/lucabecci/questions-golang-API/internal/database"
	"github.com/lucabecci/questions-golang-API/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	uri := os.Getenv("DB_URI")
	_, err := database.GetInstance(uri)
	if err != nil {
		log.Panic(err.Error())
	}
	server, err := server.GetInstance()
	if err != nil {
		log.Panic(err.Error())
	}
	server.Start(port)

	//If the developer use the Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	server.Close()

}
