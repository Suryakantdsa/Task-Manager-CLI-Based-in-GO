package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID   int
	Name string
	Done bool
}

var tasks []Task
var taskCount int = 1

func SaveTasksToFile() {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Error while saving task", err)
		return
	}
	defer file.Close()

	encoded := json.NewEncoder(file)
	err = encoded.Encode(tasks)
	if err != nil {
		fmt.Println("error while encodeing task to file", err)
		return
	}

}

func LoadTasksFromFile() {
	file, err := os.Open("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No file exist..")
			return
		}
		fmt.Println("Error while loading task..!")

	}
	defer file.Close()

	decoded := json.NewDecoder(file)
	err = decoded.Decode(&tasks)
	if err != nil {
		fmt.Println("Error while decoding the file", err)
		return
	}

}

func displayMenu() {
	fmt.Println("\n Task manager")
	fmt.Println("1. Add task")
	fmt.Println("2. View task")
	fmt.Println("3. mark as Done")
	fmt.Println("4. Delete task")
	fmt.Println("5. Exit")
	fmt.Println("Choose an option:")
}
func AddTask() {
	// var name string
	// fmt.Println("Enter name of task:")
	// fmt.Scanln(&name)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task name: ")

	// Read the entire line including spaces
	taskName, _ := reader.ReadString('\n')
	fmt.Println("task=", taskName)
	taskName = strings.TrimSpace(taskName) // Trim trailing newline or spaces
	if taskName == "" {
		fmt.Println("Task name cannot be empty.")
		return
	}
	task := Task{ID: taskCount, Name: taskName, Done: false}
	tasks = append(tasks, task)
	taskCount++
	fmt.Println("Task added sucessfully")

}
func ViewTask() {
	if len(tasks) == 0 {
		fmt.Println("No task is available")
		return
	}
	fmt.Println("\n Task:")
	for _, task := range tasks {
		status := "Pending"
		if task.Done {
			status = "Done"
		}
		fmt.Printf("%d %s [%s]\n", task.ID, task.Name, status)
	}
}
func MarkAsDone() {
	var id int
	fmt.Println("Enter task ID to mark as Done: ")
	fmt.Scanln(&id)
	for idx, task := range tasks {
		if task.ID == id {
			tasks[idx].Done = true
			fmt.Println("Task marked as Done !")
			return

		}
	}
	fmt.Println("No task found..!")

}
func DeleteTask() {
	var id int
	fmt.Println("Enter task ID to delete :")
	fmt.Scanln(&id)

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted sucessfully..!")
			return

		}
	}
	fmt.Println("No task found..!")

}

func main() {
	LoadTasksFromFile()
	for {
		displayMenu()
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			AddTask()
			SaveTasksToFile()
		case 2:
			ViewTask()
			SaveTasksToFile()
		case 3:
			MarkAsDone()
			SaveTasksToFile()
		case 4:
			DeleteTask()
			SaveTasksToFile()
		case 5:
			fmt.Println("Exiting... Goodbye.!")
			return
		default:
			fmt.Println("Invalid choice . please try again")

		}
	}

}
