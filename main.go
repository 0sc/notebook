package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

var (
	nb                                              *notebook
	doc, notification, notepad, notepadBtn, sidebar js.Value
	cb                                              js.Func
)

func newNotepad(this js.Value) {
	notepad.Set("value", "")

	notepadBtn.Call("removeEventListener", "click", cb)
	notepadBtn.Set("value", "")
	cb = callback(addNote)
	notepadBtn.Call("addEventListener", "click", cb)

}

func getNote(this js.Value) {
	id, err := strconv.Atoi(this.Get("value").String())
	if err != nil {
		notify(err.Error(), "red")
		return
	}

	note, err := nb.get(id)
	if err != nil {
		notify(err.Error(), "red")
		return
	}

	notepad.Set("value", note.content)
	notepadBtn.Call("removeEventListener", "click", cb)
	cb = callback(updateNote)
	notepadBtn.Call("addEventListener", "click", cb)
	notepadBtn.Set("value", id)
}

func addNote(_ js.Value) {
	content := notepad.Get("value").String()
	note, err := nb.add(content)
	if err != nil {
		notify(err.Error(), "red")

		return
	}

	notepadBtn.Call("removeEventListener", "click", cb)
	cb = callback(updateNote)
	notepadBtn.Call("addEventListener", "click", cb)
	notepadBtn.Set("value", nb.lastIndex)

	sidebar.Call(
		"insertBefore",
		sidebarItem(nb.lastIndex, note),
		sidebar.Get("firstChild"),
	)
	notify("note added successfully", "green")
}

func updateNote(this js.Value) {
	id, err := strconv.Atoi(this.Get("value").String())
	if err != nil {
		notify(err.Error(), "red")

		return
	}

	if !nb.exists(id) {
		addNote(this)
		return
	}

	content := notepad.Get("value").String()
	note, err := nb.update(id, content)
	if err != nil {
		notify(err.Error(), "red")

		return
	}

	sidebar.Call("replaceChild", sidebarItem(id, note), elemByID(id))
	notify("note updated successfully", "green")
}

func deleteNote(this js.Value) {
	id, err := strconv.Atoi(this.Get("value").String())
	if err != nil {
		notify(err.Error(), "red")

		return
	}

	err = nb.delete(id)
	if err != nil {
		notify(err.Error(), "red")

		return
	}

	sidebar.Call("removeChild", elemByID(id))
	notify("note deleted successfully", "green")
}

func main() {
	nb = &notebook{
		notes: make(map[int]*note),
	}

	doc = js.Global().Get("document")
	notification = elemByID("notification")
	notepad = elemByID("notepad")
	notepadBtn = elemByID("save-btn")
	sidebar = elemByID("notes")

	registerCallbacks()
	newNotepad(js.Value{})

	fmt.Println("WASM Go Initialized!!!!!")
	select {}
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(
		func(this js.Value, values []js.Value) interface{} {
			newNotepad(this)
			return nil
		},
	))

	js.Global().Set("addNote", js.FuncOf(
		func(this js.Value, values []js.Value) interface{} {
			addNote(this)
			return nil
		},
	))

	js.Global().Set("getNote", js.FuncOf(
		func(this js.Value, values []js.Value) interface{} {
			getNote(values[0])
			return nil
		},
	))

	js.Global().Set("updateNote", js.FuncOf(
		func(this js.Value, values []js.Value) interface{} {
			updateNote(values[0])
			return nil
		},
	))

	js.Global().Set("deleteNote", js.FuncOf(
		func(this js.Value, values []js.Value) interface{} {
			deleteNote(values[0])
			return nil
		},
	))
}
