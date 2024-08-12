package todo_list

import (
	"bufio"
	"log"
	"os"

	"github.com/jpeterburs/todo-txt/internal/task"
)

func ReadTasksFromFile(path string) []task.Task {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	tasks := make([]task.Task, 0)
	for scanner.Scan() {
		tasks = append(tasks, task.NewTask(scanner.Text()))
	}
	scannerErr := scanner.Err()
	if scannerErr != nil {
		log.Fatal(err)
	}

	return tasks
}
