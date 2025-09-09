package main

import (
	"os"
	"fmt"
	"todo/internal/utils"
	"strings"
)

const dataFilePath = "./data/ToDo.json"

func cli(data map[string]interface{}) map[string]interface{} {
	println("Enter command (add/remove/list):")
	var command string
	fmt.Scanln(&command)

	switch strings.ToLower(command) {
	case "add":
		println("Enter task to add:")
		var task string
		fmt.Scanln(&task)
		println("Enter priority (low, medium, high):")
		var priority string
		fmt.Scanln(&priority)
		err := utils.AddTask(data, task, priority)
		if err != nil {
			println("Error adding task: " + err.Error())
		}
	case "remove":
		println("Enter task to remove:")
		var task string
		fmt.Scanln(&task)
		err := utils.RemoveTask(data, task)
		if err != nil {
			println("Error removing task: " + err.Error())
		}
	case "list":
		list_tasks := utils.ListTasks(data)
		if len(list_tasks) == 0 {
			println("🎉Aeho! No tasks!")
		}
		for i, task := range list_tasks {
			row := fmt.Sprintf("- %d: %s", i+1, task)
			println(row)
		}
	}
	return data
}

func main() {
		data, err := utils.ReadJSONFile(dataFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			// interface{} is a type for any value
			// So map is a hash from string -> any value
			// make allocates it
			data = make(map[string]interface{})
			_ = utils.WriteJSONFile(dataFilePath, data)
			println("Created new data file at " + dataFilePath)
		} else {
			println("Error reading file: " + err.Error())
			return
		}
	}

	data = cli(data)

	_ = utils.WriteJSONFile(dataFilePath, data)
}
