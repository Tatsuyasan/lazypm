// Package list provides a customizable list view component using the Bubble Tea framework.
package list

import (
	"github.com/charmbracelet/bubbles/list"
)

func NewListView(items []customItem) list.Model {
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = item
	}

	delegate := customDelegate{}
	l := list.New(listItems, delegate, 0, 0)
	l.SetShowTitle(false)

	return l
}
