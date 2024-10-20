package mylib

import (
	"encoding/json"
  "bufio"
	"errors"
	"fmt"
	"os"
)

type Item struct {
	Task      string
	Completed bool
}

type List []Item

func (l *List) AddTask(task string) {
	new_task := Item{
		Task:      task,
		Completed: false,
	}

	*l = append(*l, new_task)
}

func (l *List) DeleteTask(index int) error {
	lp := *l
	if index <= 0 || index > len(lp) {
		return fmt.Errorf("Item %d does not exist", index)
	}

	*l = append(lp[:index-1], lp[index:]...)

	return nil
}

func (l *List) CompleteTask(index int) error {
	lp := *l
	if index <= 0 || index > len(lp) {
		return fmt.Errorf("Item %d does not exist", index)
	}

	lp[index-1].Completed = true

	return nil
}

func (l *List) DisplayList() {
	lp := *l

	fmt.Println("\tTodo-List")
  fmt.Println(" S/N|\tTask")
  fmt.Println("----+--------------------------------------------------------------------")

	switch {
	case len(lp) == 0:
		fmt.Println("\tTodolist is empty")

	case len(lp) > 0:
		for index, item := range lp {

			fmt.Printf("%3d.| %s", index + 1, item.Task)
      if item.Completed {
        fmt.Print("  ✔ \n")
      } else {
        fmt.Println()
      } 

		}
	}
}

func (l *List) SaveList(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

func (l *List) GetList(filename string) error {
	file, err := os.ReadFile(filename)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

 
func Input(quote string) string {
	fmt.Println(quote)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()

	if err != nil {
    fmt.Fprint(os.Stderr, err)
	}

	return scanner.Text()
}
