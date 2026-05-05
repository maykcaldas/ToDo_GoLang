package main

import (
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	id          int
	title       string
	description string
	status      string
}

func add(tasks []Task, task Task) []Task {
	tasks = append(tasks, task)
	list(tasks)
	return tasks
}

func list(tasks []Task) {
	for _, task := range tasks {
		fmt.Println("- ", task.String())
	}
}

func (t Task) String() string {
	return fmt.Sprintf("\tID: %d,\n\tTitle: %s,\n\tDescription: %s,\n\tStatus: %s", t.id, t.title, t.description, t.status)
}

func (t *Task) Complete() {
	t.status = "complete"
}

func markComplete(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.id == id {
			tasks[i].Complete()
			break
		}
	}
	list(tasks)
	return tasks
}

func help() {
	fmt.Println("Available commands:")
	fmt.Println("  add <task>  - Add task to the manager")
	fmt.Println("  list        - List tasks in the manager")
	fmt.Println("  help        - Show this help message")
}

func main() {
	tasks := []Task{
		{id: 1, title: "Task 1", description: "Description of Task 1", status: "pending"},
		{id: 2, title: "Task 2", description: "Description of Task 2", status: "pending"},
	}
	lastIDCreated := len(tasks)

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input>")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Title is required for adding a task")
			return
		}
		title := os.Args[2]
		description := ""
		if len(os.Args) > 3 {
			description = os.Args[3]
		}
		task := Task{
			id:          lastIDCreated + 1,
			title:       title,
			description: description,
			status:      "pending",
		}

		fmt.Println("Adding item...")
		tasks = add(tasks, task)
		lastIDCreated++
	case "list":
		fmt.Println("Listing items...")
		list(tasks)
	case "help":
		help()
	case "markComplete":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		tasks = markComplete(tasks, id)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		help()
	}
}
