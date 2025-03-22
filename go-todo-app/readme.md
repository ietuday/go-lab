Now, let's **convert this CLI-based To-Do app into a REST API** using **Go Fiber** (a lightweight web framework similar to Express.js). 🚀  

---

## **Step 1: Install Go Fiber**
First, install the **Go Fiber** package:  
```sh
go get github.com/gofiber/fiber/v2
```

---

## **Step 2: Project Structure**
Create a project structure like this:  
```
go-todo-api/
│── main.go          # Main entry point
│── handlers.go      # API logic
│── storage.go       # File/database operations
│── go.mod           # Go module dependencies
```

---

## **Step 3: Define the Task Model**
Modify `main.go` to define the **Task struct** and initialize Fiber:
```go
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
	app := fiber.New() // Create a new Fiber app

	// Routes
	app.Get("/tasks", listTasks)
	app.Post("/tasks", addTask)
	app.Put("/tasks/:id/done", markDone)
	app.Delete("/tasks/:id", deleteTask)

	app.Listen(":3000") // Start the server on port 3000
}
```

---

## **Step 4: Implement API Handlers (`handlers.go`)**
Create a new file **`handlers.go`** for API logic.

### **1. Get All Tasks**
```go
func listTasks(c *fiber.Ctx) error {
	return c.JSON(tasks)
}
```
📌 **Returns a JSON list of all tasks.**

---

### **2. Add a New Task**
```go
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
```
📌 **Handles adding new tasks with JSON input validation.**

---

### **3. Mark Task as Done**
```go
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
```
📌 **Marks a task as complete using `PUT /tasks/:id/done`.**

---

### **4. Delete a Task**
```go
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
```
📌 **Deletes a task using `DELETE /tasks/:id`.**

---

## **Step 5: Run the API**
Run the server:
```sh
go run main.go
```
📌 The API runs on `http://localhost:3000`.

---

## **Step 6: Test with Postman or `curl`**
### **1. Add a Task**
```sh
curl -X POST http://localhost:3000/tasks -H "Content-Type: application/json" -d '{"name": "Learn Go Fiber"}'
```
📌 Response:
```json
{
  "id": 1,
  "name": "Learn Go Fiber",
  "done": false
}
```

### **2. Get All Tasks**
```sh
curl http://localhost:3000/tasks
```

### **3. Mark Task as Done**
```sh
curl -X PUT http://localhost:3000/tasks/1/done
```

### **4. Delete a Task**
```sh
curl -X DELETE http://localhost:3000/tasks/1
```

---