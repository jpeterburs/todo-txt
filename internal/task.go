package task

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type task struct {
	completed      bool
	priority       string
	completionDate *time.Time
	creationDate   *time.Time
	description    string
	project        string
	context        string
}

func newTask(input string) task {
	cleanedInput := input

	completed := false
	if input[0] == 'x' {
		completed = true
		cleanedInput = input[2:]
	}

	var completionDate, creationDate *time.Time
	dateRe := regexp.MustCompile(`\d{4}\-\d{2}\-\d{2}`)
	dateMatches := dateRe.FindAllString(cleanedInput, -1)
	if len(dateMatches) > 0 {
		if completed {
			completionDate = parseDate(dateMatches[0])

			if len(dateMatches) > 1 {
				creationDate = parseDate(dateMatches[1])
			}
		} else if len(dateMatches) > 0 {
			creationDate = parseDate(dateMatches[0])
		}
	}
	cleanedInput = dateRe.ReplaceAllLiteralString(cleanedInput, "")

	prioRe := regexp.MustCompile(`\((.)\)`)
	prioMatch := prioRe.FindStringSubmatch(cleanedInput)
	var priority string
	if len(prioMatch) == 2 {
		priority = prioMatch[1]
		cleanedInput = prioRe.ReplaceAllString(cleanedInput, "")
	}

	projectRe := regexp.MustCompile(`\+\S+`)
	projectMatch := projectRe.FindString(cleanedInput)
	var project string
	if projectMatch != "" {
		project = projectMatch[1:]
		cleanedInput = projectRe.ReplaceAllLiteralString(cleanedInput, "")
	}

	contextRe := regexp.MustCompile(`\@\S+`)
	contextMatch := contextRe.FindString(cleanedInput)
	var context string
	if contextMatch != "" {
		context = contextMatch[1:]
		cleanedInput = contextRe.ReplaceAllLiteralString(cleanedInput, "")
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
		context:        context,
	}
}

func parseDate(dateStr string) *time.Time {
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return nil
	}

	return &parsedDate
}
