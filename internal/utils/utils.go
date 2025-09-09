package utils

// Todo: separate this utils in io and data utils.
// But the compiler is mad at me when I try to do that now.

import (
	"encoding/json"
	"os"
	"fmt"
	"strings"
)

func ReadJSONFile(filename string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func WriteJSONFile(filename string, data map[string]interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		return err
	}

	return nil
}

func _validatePriority(priority string) bool {
	is_valid := false
	for _, p := range []string{"low", "medium", "high"} {
		if priority == p {
			is_valid = true
		}
	}
	return is_valid
}

func AddTask(data map[string]interface{}, task string, priority string) error {
	if strings.TrimSpace(task) == "" {
		return fmt.Errorf("task cannot be empty")
	}
	
	if !_validatePriority(priority) {
		return fmt.Errorf("invalid priority")
	}
	
	exist, ok := data[task].(interface{})
	if ok && exist != nil {
		return fmt.Errorf("task already exists")
	} else {
		data[task] = priority
	}
	return nil
}

func RemoveTask(data map[string]interface{}, task string) error {
	for t, _ := range data {
		if strings.ToLower(t) == strings.ToLower(task) {
			delete(data, t)
			return nil
		}
	}
	return fmt.Errorf("task not found")
}

func ListTasks(data map[string]interface{}) []string {
	list_tasks := []string{}
	for task, _ := range data {
		list_tasks = append(list_tasks, task)
	}
	return list_tasks
}