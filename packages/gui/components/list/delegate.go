package list

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type customDelegate struct{}

// Styles for focused and normal states
var (
	itemStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240")) // Discrete color
	focusedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255")) // Highlighted color
)

func (d customDelegate) Height() int {
	return 1 // Single line per item
}

func (d customDelegate) Spacing() int {
	return 0
}

func (d customDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d customDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	i, ok := item.(customItem)
	if !ok {
		return
	}

	// Apply different styles based on focus
	style := itemStyle
	if index == m.Index() {
		style = focusedItemStyle
	}

	// Render name and description on the same line
	fmt.Fprintf(w, style.Render(i.Name+": "+i.Description))
}

// customItem represents a single item in the list with a name and description
type customItem struct {
	Name        string
	Description string
}

func (i customItem) FilterValue() string {
	return i.Name
}
