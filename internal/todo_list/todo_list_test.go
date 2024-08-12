package todo_list

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jpeterburs/todo-txt/internal/task"
	"github.com/stretchr/testify/assert"
)

// nil *time.Time pointer
var nilTime *time.Time

func TestReadTasksFromFile(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	tasks := ReadTasksFromFile(filepath.Join(wd, "..", "..", "testdata", "todo.txt"))

	expectedTasks := []task.Task{
		{
			Completed:      false,
			Priority:       "B",
			CompletionDate: nilTime,
			CreationDate:   nilTime,
			Description:    "get flowers",
			Project:        "",
			Context:        "",
		},
		{
			Completed:      false,
			Priority:       "A",
			CompletionDate: nilTime,
			CreationDate:   nilTime,
			Description:    "finish todo app",
			Project:        "",
			Context:        "",
		},
		{
			Completed:      false,
			Priority:       "",
			CompletionDate: nilTime,
			CreationDate:   nilTime,
			Description:    "repair car",
			Project:        "",
			Context:        "",
		},
	}

	assert.ElementsMatch(t, expectedTasks, tasks)
}
