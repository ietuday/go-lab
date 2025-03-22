Sure! Let's go through each part of the code step by step.  

---

## **Code Breakdown**
### **1. Importing Required Packages**
```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)
```
- **`package main`** → Defines the main package, which is the entry point of a Go program.  
- **`import` block** → Imports necessary Go standard libraries:
  - `encoding/json` → Handles JSON encoding/decoding.
  - `fmt` → Provides functions for formatted I/O (printing to console, reading input).
  - `os` → Handles file operations (reading, writing).

---

### **2. Defining the Task Structure**
```go
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}
```
- **`type Task struct {}`** → Defines a struct named `Task` to represent a to-do item.  
- **Fields in the struct:**
  - `ID int` → Unique task ID.
  - `Name string` → Name of the task.
  - `Done bool` → `true` if the task is completed, `false` otherwise.
- **JSON Tags (`json:"id"`)** → Used to specify how the struct fields will be serialized to JSON.

---

### **3. Global Variables**
```go
var tasks []Task
var filename = "tasks.json"
```
- **`tasks []Task`** → Declares a slice (dynamic array) to store all tasks.  
- **`filename = "tasks.json"`** → The file where tasks will be saved and loaded.

---

### **4. Main Function (Entry Point)**
```go
func main() {
	loadTasks() // Load tasks from the file at startup
	for {
		fmt.Println("\n1. Add Task\n2. List Tasks\n3. Mark Done\n4. Exit")
		fmt.Print("Choose an option: ")
		var choice int
		fmt.Scanln(&choice) // Read user input

		switch choice {
		case 1:
			addTask()
		case 2:
			listTasks()
		case 3:
			markDone()
		case 4:
			saveTasks() // Save tasks before exiting
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, try again!")
		}
	}
}
```
- **`loadTasks()`** → Reads existing tasks from `tasks.json` at startup.  
- **Infinite `for` loop** → Keeps the program running until the user chooses to exit.  
- **Menu Display (`fmt.Println`)** → Shows available options.  
- **User Input (`fmt.Scanln(&choice)`)** → Reads the user’s selection.  
- **`switch` Statement** → Executes the corresponding function based on the user’s choice.

---

### **5. Function to Add a Task**
```go
func addTask() {
	fmt.Print("Enter task name: ")
	var name string
	fmt.Scanln(&name) // Read user input

	task := Task{ID: len(tasks) + 1, Name: name, Done: false} // Create new task
	tasks = append(tasks, task) // Add task to the list
	saveTasks() // Save tasks to file
	fmt.Println("Task added successfully!")
}
```
- **Prompts the user (`fmt.Print`)** to enter a task name.  
- **Reads input (`fmt.Scanln(&name)`)** and stores it in `name`.  
- **Creates a new `Task` struct** and assigns an ID based on the current length of `tasks`.  
- **`append(tasks, task)`** → Adds the new task to the list.  
- **`saveTasks()`** → Saves the updated task list to `tasks.json`.

---

### **6. Function to List All Tasks**
```go
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
```
- **Iterates (`for _, task := range tasks`)** over the tasks list.  
- **Uses `if task.Done`** to check if the task is completed.  
- **Prints each task (`fmt.Printf`)** with an emoji:
  - `❌` → If task is not completed.
  - `✅` → If task is completed.

---

### **7. Function to Mark a Task as Done**
```go
func markDone() {
	fmt.Print("Enter task ID to mark as done: ")
	var id int
	fmt.Scanln(&id) // Read user input

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
```
- **Prompts the user (`fmt.Print`)** to enter a task ID.  
- **Iterates through the `tasks` list** to find the task with the matching ID.  
- **If found, sets `task.Done = true`** and saves the updated list.  
- **If no task matches, prints "Task not found!"**.

---

### **8. Function to Load Tasks from File**
```go
func loadTasks() {
	file, err := os.ReadFile(filename)
	if err == nil {
		json.Unmarshal(file, &tasks)
	}
}
```
- **`os.ReadFile(filename)`** → Reads the contents of `tasks.json`.  
- **If the file exists, `json.Unmarshal(file, &tasks)`** converts JSON data into a Go struct.  
- **If the file does not exist, it does nothing** (avoiding errors).  

---

### **9. Function to Save Tasks to File**
```go
func saveTasks() {
	data, _ := json.Marshal(tasks) // Convert tasks to JSON
	os.WriteFile(filename, data, 0644) // Save JSON data to file
}
```
- **`json.Marshal(tasks)`** → Converts the `tasks` slice into JSON format.  
- **`os.WriteFile(filename, data, 0644)`** → Writes the JSON data to `tasks.json`.  

---

## **How It Works in Action**
```sh
$ go run main.go

1. Add Task
2. List Tasks
3. Mark Done
4. Exit
Choose an option: 1
Enter task name: Buy groceries
Task added successfully!

1. Add Task
2. List Tasks
3. Mark Done
4. Exit
Choose an option: 2

Tasks:
1. Buy groceries [❌]

1. Add Task
2. List Tasks
3. Mark Done
4. Exit
Choose an option: 3
Enter task ID to mark as done: 1
Task marked as done!

1. Add Task
2. List Tasks
3. Mark Done
4. Exit
Choose an option: 2

Tasks:
1. Buy groceries [✅]

1. Add Task
2. List Tasks
3. Mark Done
4. Exit
Choose an option: 4
Exiting...
```
- The program saves tasks in `tasks.json`, so when restarted, tasks are not lost.

---