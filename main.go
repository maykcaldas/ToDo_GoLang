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

func (t Task) String() string {
	return fmt.Sprintf("\tID: %d,\n\tTitle: %s,\n\tDescription: %s,\n\tStatus: %s", t.id, t.title, t.description, t.status)
}

func (t *Task) Complete() {
	t.status = "complete"
}

type TaskList []Task

func (tl TaskList) String() string {
	result := ""
	for _, task := range tl {
		result += task.String() + "\n"
	}
	return result
}

func (tl *TaskList) Add(task Task) TaskList {
	*tl = append(*tl, task)
	return *tl
}

func (tl *TaskList) MarkComplete(id int) {
	for i, task := range *tl {
		if task.id == id {
			(*tl)[i].Complete()
			break
		}
	}
}

func help() {
	fmt.Println("Available commands:")
	fmt.Println("  add <task>  - Add task to the manager")
	fmt.Println("  list        - List tasks in the manager")
	fmt.Println("  help        - Show this help message")
}

func main() {
	tasks := TaskList{
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
		tasks = tasks.Add(task)
		fmt.Println(tasks)
		lastIDCreated++
	case "list":
		fmt.Println("Listing items...")
		fmt.Println(tasks)
	case "help":
		help()
	case "markComplete":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		tasks.MarkComplete(id)
		fmt.Printf("Marked task %d as complete.\n", id)
		fmt.Println(tasks)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		help()
	}
}
