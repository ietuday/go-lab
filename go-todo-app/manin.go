package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var tasks []Task
var filename = "tasks.json"

func main() {
	loadTasks() // Load existing tasks

	for {
		fmt.Println("\n1. Add Task\n2. List Tasks\n3. Mark Done\n4. Delete Task\n5. Exit")
		fmt.Print("Choose an option: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask()
		case 2:
			listTasks()
		case 3:
			markDone()
		case 4:
			deleteTask() // New function to delete a task
		case 5:
			saveTasks()
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, try again!")
		}
	}
}

func addTask() {
	fmt.Print("Enter task name: ")
	var name string
	fmt.Scanln(&name)

	task := Task{ID: len(tasks) + 1, Name: name, Done: false}
	tasks = append(tasks, task)
	saveTasks()
	fmt.Println("Task added successfully!")
}

func listTasks() {
	fmt.Println("\nTasks:")
	for _, task := range tasks {
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("%d. %s [%s]\n", task.ID, task.Name, status)
	}
}

func markDone() {
	fmt.Print("Enter task ID to mark as done: ")
	var id int
	fmt.Scanln(&id)

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			saveTasks()
			fmt.Println("Task marked as done!")
			return
		}
	}
	fmt.Println("Task not found!")
}

func loadTasks() {
	file, err := os.ReadFile(filename)
	if err == nil {
		json.Unmarshal(file, &tasks)
	}
}

func saveTasks() {
	data, _ := json.Marshal(tasks)
	os.WriteFile(filename, data, 0644)
}

func deleteTask() {
	fmt.Print("Enter task ID to delete: ")
	var id int
	fmt.Scanln(&id)

	index := -1
	for i, task := range tasks {
		if task.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Task not found!")
		return
	}

	// Remove task from slice
	tasks = append(tasks[:index], tasks[index+1:]...)
	saveTasks()
	fmt.Println("Task deleted successfully!")
}
