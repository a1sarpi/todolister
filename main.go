package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID   int
	Text string
	Done bool
}

var tasks []Task
var currentID = 1

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nTo-Do List Manager")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("	a. Show Incomplete Tasks")
		fmt.Println("	b. Show Complete Tasks")
		fmt.Println("	c: Show All (default)")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			addTask(reader)
		case "2":
			fmt.Println("  a. Show Incomplete Tasks")
			fmt.Println("  b. Show Complete Tasks")
			fmt.Println("  c. Show All (default)")
			fmt.Print("Choose an option: ")

			listInput, _ := reader.ReadString('\n')
			listInput = strings.TrimSpace(listInput)

			switch listInput {
			case "a":
				listTasks(getIncompleteTasks())
			case "b":
				listTasks(getCompleteTasks())
			case "c":
				listTasks(tasks)
			default:
				listTasks(tasks)
			}
		case "3":
			markDone(reader)
		case "4":
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid option!")
		}
	}
}

func addTask(reader *bufio.Reader) {
	fmt.Print("Enter task: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	tasks = append(tasks, Task{
		ID:   currentID,
		Text: text,
		Done: false,
	})
	currentID++

	fmt.Println("Task added!")
}

func listTasks(t []Task) {
	if len(t) == 0 {
		fmt.Println("No tasks yet!")
		return
	}

	for _, task := range t {
		status := " "
		if task.Done {
			status = "✓"
		}

		fmt.Printf("[%d] [%s] %s\n", task.ID, status, task.Text)
	}
}

func markDone(reader *bufio.Reader) {
	listTasks(tasks)
	if len(tasks) == 0 {
		return
	}

	fmt.Print("Enter task ID to mark as done: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var id int
	_, err := fmt.Sscanf(input, "%d", &id)
	if err != nil {
		fmt.Println("Invalid ID!")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			fmt.Println("Task marked as done!")
			return
		}
	}

	fmt.Println("Task not found!")
}

func getIncompleteTasks() []Task {
	var result []Task
	for _, task := range tasks {
		if !task.Done {
			result = append(result, task)
		}
	}

	return result
}

func getCompleteTasks() []Task {
	var result []Task
	for _, task := range tasks {
		if task.Done {
			result = append(result, task)
		}
	}

	return result
}
