package main

import (
	"fmt"
)

type Task struct {
	ID   int
	Name string
	Done bool
}

var tasks []Task
var taskCount int = 1

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
	var name string
	fmt.Println("Enter name of task:")
	fmt.Scanln(&name)
	task := Task{ID: taskCount, Name: name, Done: false}
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
	for {
		displayMenu()
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			AddTask()
		case 2:
			ViewTask()
		case 3:
			MarkAsDone()
		case 4:
			DeleteTask()
		case 5:
			fmt.Println("Exiting... Goodbye.!")
			return
		default:
			fmt.Println("Invalid choice . please try again")

		}
	}

}
