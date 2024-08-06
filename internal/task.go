package task

import (
	"regexp"
	"strings"
)

type task struct {
	completed      bool
	priority       string
	completionDate string
	creationDate   string
	description    string
}

func newTask(input string) task {
	cleanedInput := input

	completed := false
	if input[0] == 'x' {
		completed = true
		cleanedInput = input[2:]
	}

	var completionDate string
	if completed {
		completionDateRe := regexp.MustCompile(`\d{4}\-\d{2}\-\d{2}\s`)
		completionDateMatch := completionDateRe.FindStringSubmatch(cleanedInput)
		if len(completionDateMatch) >= 1 {
			completionDate = strings.Join(completionDateMatch, "")[:10]
			cleanedInput = cleanedInput[11:]
		}
	}

	var creationDate string
	creationDateRe := regexp.MustCompile(`\d{4}\-\d{2}\-\d{2}\s`)
	creationDateMatch := creationDateRe.FindStringSubmatch(cleanedInput)
	if len(creationDateMatch) >= 1 {
		creationDate = strings.Join(creationDateMatch, "")[:10]
		cleanedInput = cleanedInput[11:]
	}

	prioRe := regexp.MustCompile(`\((.)\)\s`)
	prioMatch := prioRe.FindStringSubmatch(cleanedInput)
	var priority string
	if len(prioMatch) > 1 {
		priority = prioMatch[1]
		cleanedInput = prioRe.ReplaceAllString(cleanedInput, "")
	}

	return task{
		completed:      completed,
		priority:       priority,
		completionDate: completionDate,
		creationDate:   creationDate,
		description:    cleanedInput,
	}
}
