package mylib_test

import (
	"os"
	"testing"
	"todoList"
)

func TestAddTask(t *testing.T) {
  testList := mylib.List{}
  
  task := "First Task"

  testList.AddTask(task)

  if testList[0].Task != task {
    t.Errorf("Expected %s got %s",task,  testList[0].Task)
  }

  if testList[0].Completed != false {
    t.Errorf("Task cannot be completed")
  }
}

func TestCompleteTask(t *testing.T) {
  testList := mylib.List{
    {Task: "New Task", Completed: false},
  }

  testList.CompleteTask(1)

  if !testList[0].Completed {
    t.Errorf("Task should be completed")
  }
}

func TestDeleteTask(t *testing.T) {
  testList := mylib.List{}

  tasks := []string {
    "First Task",
    "Second Task",
  }

  for _, k := range(tasks) {
    testList.AddTask(k)
  }

  if testList[0].Task != tasks[0] {
    t.Errorf("Expected %s got %s", tasks[0], testList[0].Task)
  }

  testList.DeleteTask(1)

  if testList[0].Task != tasks[1] {
    t.Errorf("Expected %s got %s", tasks[1], testList[0].Task)
  }

  if len(testList) != 1 {
    t.Errorf("Expected list length of %d, got %d", 1, len(testList))
  }
}

func TestSaveGetList(t *testing.T) {
  l1 := mylib.List{
    {Task: "New Task", Completed: false},
  }
  l2 := mylib.List{}

  tf, err := os.CreateTemp("", "sample")

  if err != nil {
    t.Fatalf("Error while creating temp file: %s", err)
  }

  defer os.Remove(tf.Name())

  if err := l1.SaveList(tf.Name()); err != nil {
    t.Fatalf("Error while saving to file: %s", err)
  }

  if err := l2.GetList(tf.Name()); err != nil {
    t.Fatalf("Error while getting file: %s", err)
  }

  if l1[0].Task != l2[0].Task {
    t.Errorf("Expected %s, got %s", l1[0].Task, l2[0].Task)
  }
}
