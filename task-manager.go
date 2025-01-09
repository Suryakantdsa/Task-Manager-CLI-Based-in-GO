package main

import "fmt"

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

func main() {

}
