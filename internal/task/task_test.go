package task

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	type testCases struct {
		description    string
		input          string
		expectedStruct Task
	}

	// nil *time.Time pointer
	var nilTime *time.Time

	for _, scenario := range []testCases{
		{
			description: "a normal task",
			input:       "write tests",
			expectedStruct: Task{
				Completed:      false,
				Priority:       "",
				CompletionDate: nilTime,
				CreationDate:   nilTime,
				Description:    "write tests",
				Project:        "",
				Context:        "",
			},
		},
		{
			description: "a task with a project",
			input:       "finish task struct +todo-txt",
			expectedStruct: Task{
				Completed:      false,
				Priority:       "",
				CompletionDate: nilTime,
				CreationDate:   nilTime,
				Description:    "finish task struct",
				Project:        "todo-txt",
				Context:        "",
			},
		},
		{
			description: "a task with context",
			input:       "order pizza @pizza-place",
			expectedStruct: Task{
				Completed:      false,
				Priority:       "",
				CompletionDate: nilTime,
				CreationDate:   nilTime,
				Description:    "order pizza",
				Project:        "",
				Context:        "pizza-place",
			},
		},
		{
			description: "a task with priority",
			input:       "(A) implement functions",
			expectedStruct: Task{
				Completed:      false,
				Priority:       "A",
				CompletionDate: nilTime,
				CreationDate:   nilTime,
				Description:    "implement functions",
				Project:        "",
				Context:        "",
			},
		},
		{
			description: "a completed task",
			input:       "x 2024-08-06 write another test",
			expectedStruct: Task{
				Completed:      true,
				Priority:       "",
				CompletionDate: parseDate("2024-08-06"),
				CreationDate:   nilTime,
				Description:    "write another test",
				Project:        "",
				Context:        "",
			},
		},
		{
			description: "a completed task with priority",
			input:       "x 2024-08-06 (B) call mom",
			expectedStruct: Task{
				Completed:      true,
				Priority:       "B",
				CompletionDate: parseDate("2024-08-06"),
				CreationDate:   nilTime,
				Description:    "call mom",
				Project:        "",
				Context:        "",
			},
		},
		{
			description: "a task with creation date",
			input:       "2024-01-01 make new year resolution",
			expectedStruct: Task{
				Completed:      false,
				Priority:       "",
				CompletionDate: nilTime,
				CreationDate:   parseDate("2024-01-01"),
				Description:    "make new year resolution",
				Project:        "",
				Context:        "",
			},
		},
		{
			description: "a task with creation date and priority",
			input:       "2024-02-01 (A) bake cake",
			expectedStruct: Task{
				Completed:      false,
				Priority:       "A",
				CompletionDate: nilTime,
				CreationDate:   parseDate("2024-02-01"),
				Description:    "bake cake",
				Project:        "",
				Context:        "",
			},
		},
		{
			description: "a completed task with creation date",
			input:       "x 2024-03-31 2024-01-01 hide easter eggs for kids",
			expectedStruct: Task{
				Completed:      true,
				Priority:       "",
				CompletionDate: parseDate("2024-03-31"),
				CreationDate:   parseDate("2024-01-01"),
				Description:    "hide easter eggs for kids",
				Project:        "",
				Context:        "",
			},
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			assert.Equal(t, scenario.expectedStruct, NewTask(scenario.input))
		})
	}
}
