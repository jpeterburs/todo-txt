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
	project        string
}

func newTask(input string) task {
	cleanedInput := input

	completed := false
	if input[0] == 'x' {
		completed = true
		cleanedInput = input[2:]
	}

	var completionDate, creationDate string
	dateRe := regexp.MustCompile(`\d{4}\-\d{2}\-\d{2}`)
	dateMatches := dateRe.FindAllString(cleanedInput, -1)
	if len(dateMatches) > 0 {
		if completed {
			completionDate = dateMatches[0]

			if len(dateMatches) > 1 {
				creationDate = dateMatches[1]
			}
		} else if len(dateMatches) > 0 {
			creationDate = dateMatches[0]
		}
	}
	cleanedInput = dateRe.ReplaceAllLiteralString(cleanedInput, "")

	prioRe := regexp.MustCompile(`\((.)\)`)
	prioMatch := prioRe.FindStringSubmatch(cleanedInput)
	var priority string
	if len(prioMatch) > 1 {
		priority = prioMatch[1]
		cleanedInput = prioRe.ReplaceAllString(cleanedInput, "")
	}

	projectRe := regexp.MustCompile(`\+\S+`)
	projectMatch := projectRe.FindStringSubmatch(cleanedInput)
	var project string
	if len(projectMatch) == 1 {
		project = projectMatch[0][1:]
		cleanedInput = projectRe.ReplaceAllLiteralString(cleanedInput, "")
	}

	cleanedInput = strings.TrimSpace(cleanedInput)
	words := strings.Fields(cleanedInput)
	cleanedInput = strings.Join(words, " ")

	return task{
		completed:      completed,
		priority:       priority,
		completionDate: completionDate,
		creationDate:   creationDate,
		description:    cleanedInput,
		project:        project,
	}
}
