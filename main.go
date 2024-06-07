package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

var tasks = []Task{}
var nextID = 1

func AddTask(title string) {
	task := Task{ID: nextID, Title: title, Done: false}
	tasks = append(tasks, task)
	nextID++
}

func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	for _, task := range tasks {
		status := "Incomplete"
		if task.Done {
			status = "Complete"
		}
		fmt.Printf("%d: %s [%s]\n", task.ID, task.Title, status)
	}
}

func CompleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			return
		}
	}
}

func DeleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return
		}
	}

}

func displayMenu() {
	fmt.Println("Choose an option:")
	fmt.Println("1. Add Task")
	fmt.Println("2. List Tasks")
	fmt.Println("3. Complete Task")
	fmt.Println("4. Delete Task")
	fmt.Println("5. Exit")
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		displayMenu()
		fmt.Print("Enter choice: ")
		input, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(input))

		switch choice {
		case 1:
			fmt.Print("Enter task title: ")
			title, _ := reader.ReadString('\n')
			AddTask(strings.TrimSpace(title))
		case 2:
			ListTasks()
		case 3:
			fmt.Print("Enter task ID to complete: ")
			idInput, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idInput))
			CompleteTask(id)
		case 4:
			fmt.Print("Enter task ID to delete: ")
			idInput, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idInput))
			DeleteTask(id)
		case 5:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
