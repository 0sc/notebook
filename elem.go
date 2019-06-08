package main

import (
	"fmt"
	"syscall/js"
)

func editBtnElemFor(id int) js.Value {
	icon := doc.Call("createElement", "i")
	addClass(icon, "fas", "fa-edit")

	btn := doc.Call("createElement", "button")
	btn.Call("append", icon)
	btn.Set("value", id)
	addClass(btn, "edit-btn")
	btn.Call("addEventListener", "click", callback(getNote))

	return btn
}

func deleteBtnElemFor(id int) js.Value {
	icon := doc.Call("createElement", "i")
	addClass(icon, "far", "fa-trash-alt")

	btn := doc.Call("createElement", "button")
	btn.Call("append", icon)
	btn.Set("value", id)
	addClass(btn, "delete-btn")
	btn.Call("addEventListener", "click", callback(deleteNote))

	return btn
}

func elemByID(id interface{}) js.Value {
	return doc.Call("getElementById", id)
}

func sidebarItem(id int, n *note) js.Value {
	editBtn := editBtnElemFor(id)
	deleteBtn := deleteBtnElemFor(id)

	div := doc.Call("createElement", "div")
	addClass(div, "note-heading")
	div.Set("innerHTML", n.heading)

	c := doc.Call("createElement", "p")
	addClass(c, "note-crumbs")
	t := n.editedAt.Format("03:04PM 02.01.2006")
	c.Set("innerHTML", fmt.Sprintf("last edit: %s", t))

	div.Call("append", c)

	li := doc.Call("createElement", "li")
	li.Set("id", id)
	li.Call("append", div)
	li.Call("append", editBtn)
	li.Call("append", deleteBtn)

	return li
}

func callback(fn func(this js.Value)) js.Func {
	return js.FuncOf(func(this js.Value, _ []js.Value) interface{} {
		fn(this)
		return nil
	})
}

func notify(msg, color string) {
	notification.Set("innerHTML", msg)
	addClass(notification, color)
}

func addClass(el js.Value, classes ...string) {
	list := el.Get("classList")
	for _, class := range classes {
		list.Call("add", class)
	}
}
