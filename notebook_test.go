package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_notebook_exists(t *testing.T) {
	t.Parallel()

	nb := &notebook{notes: map[int]*note{50: &note{}}}
	tests := []struct {
		name  string
		index int
		want  bool
	}{
		{
			name:  "it returns true if index exists",
			index: 50,
			want:  true,
		},
		{
			name:  "it returns false if index does not exist",
			index: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nb.exists(tt.index)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_notebook_get(t *testing.T) {
	t.Parallel()

	n := &note{}
	nb := &notebook{notes: map[int]*note{50: n}}
	tests := []struct {
		name    string
		index   int
		want    *note
		wantErr error
	}{
		{
			name:  "it returns the note if it exists",
			index: 50,
			want:  n,
		},
		{
			name:    "it returns error if note is not found",
			index:   2,
			wantErr: errNoteNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nb.get(tt.index)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_notebook_add(t *testing.T) {
	t.Parallel()

	nb := &notebook{notes: map[int]*note{}}
	tests := []struct {
		name    string
		content string
		wantErr bool
		l       int
	}{
		{
			name:    "it returns error if content is empty",
			wantErr: true,
		},
		{
			name:    "it returns the new note",
			content: "my note",
			l:       1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := nb.add(tt.content)
			assert.Equal(t, tt.l, len(nb.notes))
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_notebook_update(t *testing.T) {
	t.Parallel()

	n := &note{}
	nb := &notebook{notes: map[int]*note{50: n}}
	tests := []struct {
		name    string
		index   int
		content string
		wantErr bool
	}{
		{
			name:    "it returns error if note does not exist",
			content: "note update",
			wantErr: true,
		},
		{
			name:    "it returns error if new content is empty",
			index:   40,
			wantErr: true,
		},
		{
			name:    "it updates note successfully",
			content: "note update",
			index:   50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := nb.update(tt.index, tt.content)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_notebook_delete(t *testing.T) {
	t.Parallel()

	nb := &notebook{notes: map[int]*note{50: &note{}}}
	tests := []struct {
		name    string
		index   int
		l       int
		wantErr bool
	}{
		{
			name:    "it returns error is note does not exist",
			index:   20,
			l:       1,
			wantErr: true,
		},
		{
			name:  "it deletes note with no errors",
			index: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := nb.delete(tt.index)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.l, len(nb.notes))
		})
	}
}
