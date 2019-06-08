package main

import (
	"fmt"
	"strings"
	"time"
)

var (
	errEmptyNote = fmt.Errorf("note is empty")
)

type note struct {
	heading  string
	content  string
	editedAt time.Time
}

func newNote(content string) (*note, error) {
	content = strings.TrimSpace(content)

	if content == "" {
		return nil, errEmptyNote
	}

	return &note{
		content:  content,
		heading:  heading(content),
		editedAt: time.Now(),
	}, nil
}

const maxHeadingLength = 15
const suffix = " ..."
func heading(content string) string {
	content = strings.SplitN(content, "\n", 2)[0]

	if len(content) > maxHeadingLength {
		content = fmt.Sprintf("%s%s", content[0:maxHeadingLength], suffix)
	}

	return strings.Title(content)
}
