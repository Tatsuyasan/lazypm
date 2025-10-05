package list

import "github.com/charmbracelet/lipgloss"

// Styles for rendering list items
var (
	ItemStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240")) // Discrete color
	FocusedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255")) // Highlighted color
)
