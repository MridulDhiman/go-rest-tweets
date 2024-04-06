package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	
)


var tweets []Tweet



/// create slice of struct
func main() {
	app := fiber.New();

	app.Get("/", HelloController)
	app.Get("/tweet", GetAllNotes)
	app.Post("/tweet", CreateNewNote)
	app.Get("/tweet/:id", GetSingleTweet)
	app.Put("/tweet/:id", UpdateTweet)
	app.Delete("/tweet/:id", DeleteTweet)
	
	log.Fatal(app.Listen(":8080"))
}
