package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"todoList"
)


const filename = "todolist.json"

func main() {

	todo_list := &mylib.List{}

	if err := todo_list.GetList(filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	for {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		todo_list.DisplayList()

		fmt.Println("\nPress [A] to Add task, [D] to Delete task, [M] to Mark task, [S] to Save, [Q] to Quit")

		var option string
		fmt.Scanf("%s", &option)
		option = strings.ToLower(option)

		switch option {

		case "a":
			task := mylib.Input("Input task to add")
			fmt.Println(task)
			todo_list.AddTask(task)

		case "d":
			input := mylib.Input("Input index of task to delete")
			index, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println("Index should be an integer")
				continue
			}

			todo_list.DeleteTask(index)

		case "m":
			input := mylib.Input("Input index of completed task")
			index, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println("Index should be an integer")
				continue
			}

			todo_list.CompleteTask(index)

		case "s":

			todo_list.SaveList(filename)

		case "q":
			break

		}
	}

}
