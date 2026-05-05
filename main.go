package main

import (
	"fmt"
	"os"
)

func add(tasks []string, task string) []string {
	tasks = append(tasks, task)
	list(tasks)
	return tasks
}

func list(tasks []string) {
	for _, task := range tasks {
		fmt.Println("- ", task)
	}
}

func help() {
	fmt.Println("Available commands:")
	fmt.Println("  add <task>  - Add task to the manager")
	fmt.Println("  list        - List tasks in the manager")
	fmt.Println("  help        - Show this help message")
}

func main() {
	tasks := []string{"new task", "nice task"}
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input>")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		task := os.Args[2]
		fmt.Println("Adding item...")
		add(tasks, task)
	case "list":
		fmt.Println("Listing items...")
		list(tasks)
	case "help":
		help()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		help()
	}
}
