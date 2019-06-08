package main

import (
	"fmt"
)

var (
	errNoteNotFound = fmt.Errorf("note not found")
)

type notebook struct {
	notes     map[int]*note
	lastIndex int
}

func (nb *notebook) exists(index int) bool {
	_, ok := nb.notes[index]
	return ok
}

func (nb *notebook) get(index int) (*note, error) {
	note, ok := nb.notes[index]

	var err error
	if !ok {
		err = errNoteNotFound
	}

	return note, err
}

func (nb *notebook) add(content string) (*note, error) {
	n, err := newNote(content)
	if err != nil {
		return nil, err
	}

	nb.lastIndex++
	nb.notes[nb.lastIndex] = n

	return n, nil
}

func (nb *notebook) update(index int, content string) (*note, error) {
	if !nb.exists(index) {
		return nil, errNoteNotFound
	}

	n, err := newNote(content)
	if err != nil {
		return nil, err
	}

	nb.notes[index] = n

	return n, nil
}

func (nb *notebook) delete(index int) error {
	if !nb.exists(index) {
		return errNoteNotFound
	}

	delete(nb.notes, index)

	return nil
}
