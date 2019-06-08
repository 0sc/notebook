package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newNote(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		content string
		wantErr bool
	}{
		{
			name:    "it returns an error if content is empty",
			wantErr: true,
		},
		{
			name:    "it returns a new note",
			content: "something",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := newNote(tt.content)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_heading(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		content string
		wantLen int
	}{
		{
			name:    "it returns the content if less than maxHeadingLength",
			content: strings.Repeat("i", maxHeadingLength),
			wantLen: maxHeadingLength,
		},
		{
			name:    "it truncates the content if more than maxHeadingLength",
			content: strings.Repeat("i", maxHeadingLength+1),
			wantLen: maxHeadingLength + len(suffix),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := heading(tt.content)
			assert.Equal(t, tt.wantLen, len(got))
		})
	}
}
