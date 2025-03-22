package main

import (
	"github.com/gofiber/fiber/v2"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var tasks []Task

func main() {
	app := fiber.New()

	// Define API Routes
	app.Get("/tasks", listTasks)
	app.Post("/tasks", addTask)
	app.Put("/tasks/:id/done", markDone)
	app.Delete("/tasks/:id", deleteTask)

	app.Listen(":3000")
}
