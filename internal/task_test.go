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
			},
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			assert.Equal(t, scenario.expectedStruct, newTask(scenario.input))
		})
	}
}
