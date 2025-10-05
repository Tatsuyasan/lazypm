// Package scripts provides the scripts list view for the Bubble Tea-based user interface.
package scripts

import (
	"github.com/Tatsuyasan/lazyPm/packages/context"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var listBorderStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("63")).
	Padding(1).
	Width(30)

// ScriptsModel represents a simple list model
type ScriptsModel struct {
	list list.Model
}

func (m ScriptsModel) Init() tea.Cmd {
	return nil
}

func (m ScriptsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ScriptsModel) View() string {
	return listBorderStyle.Render(m.list.View())
}

func NewScriptsModel(items []context.ScriptInfo) ScriptsModel {
	listItems := make([]list.Item, len(items))
	// for i, item := range items {
	// 	listItems[i] = list.Item(item.Name)
	// }

	l := list.New(listItems, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Basic List"
	l.SetShowTitle(true)

	return ScriptsModel{list: l}
}
