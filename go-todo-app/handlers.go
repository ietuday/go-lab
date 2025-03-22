package main

import (
	"github.com/gofiber/fiber/v2"
)

func listTasks(c *fiber.Ctx) error {
	return c.JSON(tasks)
}

func addTask(c *fiber.Ctx) error {
	type Request struct {
		Name string `json:"name"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if req.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Task name cannot be empty"})
	}

	task := Task{ID: len(tasks) + 1, Name: req.Name, Done: false}
	tasks = append(tasks, task)
	return c.Status(201).JSON(task)
}

func markDone(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			return c.JSON(tasks[i])
		}
	}

	return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
}

func deleteTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	index := -1
	for i, task := range tasks {
		if task.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	return c.JSON(fiber.Map{"message": "Task deleted"})
}
