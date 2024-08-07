package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	type testCases struct {
		description    string
		input          string
		expectedStruct task
	}

	for _, scenario := range []testCases{
		{
			description: "a normal task",
			input:       "write tests",
			expectedStruct: task{
				completed:      false,
				priority:       "",
				completionDate: "",
				creationDate:   "",
				description:    "write tests",
				project:        "",
				context:        "",
			},
		},
		{
			description: "a task with a project",
			input:       "finish task struct +todo-txt",
			expectedStruct: task{
				completed:      false,
				priority:       "",
				completionDate: "",
				creationDate:   "",
				description:    "finish task struct",
				project:        "todo-txt",
				context:        "",
			},
		},
		{
			description: "a task with context",
			input:       "order pizza @pizza-place",
			expectedStruct: task{
				completed:      false,
				priority:       "",
				completionDate: "",
				creationDate:   "",
				description:    "order pizza",
				project:        "",
				context:        "pizza-place",
			},
		},
		{
			description: "a task with priority",
			input:       "(A) implement functions",
			expectedStruct: task{
				completed:      false,
				priority:       "A",
				completionDate: "",
				creationDate:   "",
				description:    "implement functions",
				project:        "",
				context:        "",
			},
		},
		{
			description: "a completed task",
			input:       "x 2024-08-06 write another test",
			expectedStruct: task{
				completed:      true,
				priority:       "",
				completionDate: "2024-08-06",
				creationDate:   "",
				description:    "write another test",
				project:        "",
				context:        "",
			},
		},
		{
			description: "a completed task with priority",
			input:       "x 2024-08-06 (B) call mom",
			expectedStruct: task{
				completed:      true,
				priority:       "B",
				completionDate: "2024-08-06",
				creationDate:   "",
				description:    "call mom",
				project:        "",
				context:        "",
			},
		},
		{
			description: "a task with creation date",
			input:       "2024-01-01 make new year resolution",
			expectedStruct: task{
				completed:      false,
				priority:       "",
				completionDate: "",
				creationDate:   "2024-01-01",
				description:    "make new year resolution",
				project:        "",
				context:        "",
			},
		},
		{
			description: "a task with creation date and priority",
			input:       "2024-02-01 (A) bake cake",
			expectedStruct: task{
				completed:      false,
				priority:       "A",
				completionDate: "",
				creationDate:   "2024-02-01",
				description:    "bake cake",
				project:        "",
				context:        "",
			},
		},
		{
			description: "a completed task with creation date",
			input:       "x 2024-03-31 2024-01-01 hide easter eggs for kids",
			expectedStruct: task{
				completed:      true,
				priority:       "",
				completionDate: "2024-03-31",
				creationDate:   "2024-01-01",
				description:    "hide easter eggs for kids",
				project:        "",
				context:        "",
			},
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			assert.Equal(t, scenario.expectedStruct, newTask(scenario.input))
		})
	}
}
