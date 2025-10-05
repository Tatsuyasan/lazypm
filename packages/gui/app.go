// Package gui provides the Bubble Tea-based user interface for lazyPm.
package gui

import (
	"github.com/Tatsuyasan/lazyPm/packages/context"
	"github.com/Tatsuyasan/lazyPm/packages/gui/views/scripts"
	tea "github.com/charmbracelet/bubbletea"
)

type MainAppModel struct {
	scriptsModel scripts.ScriptsModel
}

func (m MainAppModel) Init() tea.Cmd {
	return m.scriptsModel.Init()
}

func (m MainAppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	newModel, cmd := m.scriptsModel.Update(msg)
	m.scriptsModel = newModel.(scripts.ScriptsModel)
	return m, cmd
}

func (m MainAppModel) View() string {
	return m.scriptsModel.View()
}

func RunGui() {
	ctx := context.GetContext()

	items, _ := ctx.Manager.ListScripts()

	model := MainAppModel{
		scriptsModel: scripts.NewScriptsModel(items),
	}

	p := tea.NewProgram(model)
	p.Run()
}
