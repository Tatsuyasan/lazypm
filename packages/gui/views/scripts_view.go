package views

import (
	"fmt"

	"github.com/Tatsuyasan/lazyPm/packages/gui/components"
	"github.com/Tatsuyasan/lazyPm/packages/helpers"
	"github.com/Tatsuyasan/lazyPm/packages/models"
	"github.com/jroimartin/gocui"
)

type ScriptsView struct {
	*components.BasePanel
	manager  models.PackageManager
	commands []string
}

func NewScriptsView() *ScriptsView {
	config := components.BasePanelConfig{
		Name:        "scripts",
		Title:       "Scripts & Commands",
		Highlighted: true,
		Selectable:  true,
	}
	
	return &ScriptsView{
		BasePanel: components.NewBasePanel(config),
		commands:  make([]string, 0),
	}
}

func (sv *ScriptsView) Render(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}
	
	sv.BasePanel.Render(g, v)
	
	v.Clear()
	v.Title = sv.Config.Title
	
	if len(sv.Items) == 0 {
		if err := sv.loadScriptsAndCommands(); err != nil {
			fmt.Fprintf(v, "Error loading scripts/commands: %v", err)
			return nil
		}
	}
	
	if len(sv.Items) == 0 {
		if sv.manager != nil {
			switch sv.manager.Name() {
			case "go":
				fmt.Fprintf(v, "No commands found")
			default:
				fmt.Fprintf(v, "No scripts or commands found")
			}
		} else {
			fmt.Fprintf(v, "No scripts or commands found")
		}
		return nil
	}
	
	_, viewHeight := v.Size()
	viewHeight -= 2
	
	sv.AdjustScrollOffset(viewHeight)
	
	visibleItems := sv.GetVisibleItems(viewHeight)
	visibleSelectedIndex := sv.GetVisibleSelectedIndex()
	
	viewWidth, _ := v.Size()
	for i, item := range visibleItems {
		actualIndex := sv.ScrollOffset + i
		scriptsCount := len(sv.Items) - len(sv.commands)
		var content string
		if actualIndex < scriptsCount {
			content = fmt.Sprintf("📜 %s", item)
		} else {
			content = fmt.Sprintf("⚙️  %s", item)
		}
		padding := viewWidth - len(content)
		if padding > 0 {
			content += fmt.Sprintf("%*s", padding, "")
		}
		fmt.Fprintf(v, "%s\n", content)
	}
	
	if len(visibleItems) > 0 && visibleSelectedIndex >= 0 && visibleSelectedIndex < len(visibleItems) {
		v.SetCursor(0, visibleSelectedIndex)
	}
	
	return nil
}

func (sv *ScriptsView) HandleKey(g *gocui.Gui, v *gocui.View, key interface{}) error {
	_, viewHeight := v.Size()
	viewHeight -= 2
	
	switch key {
	case 'j':
		sv.MoveDownWithScroll(viewHeight)
	case 'k':
		sv.MoveUpWithScroll(viewHeight)
	case gocui.KeyEnter:
		return sv.Execute(g)
	}
	return nil
}

func (sv *ScriptsView) loadScriptsAndCommands() error {
	return helpers.WithManager("", func(manager models.PackageManager) error {
		sv.manager = manager
		
		var scripts []string
		var err error
		
		switch manager.Name() {
		case "npm", "pnpm", "yarn":
			scripts, err = manager.ListScripts()
			if err != nil {
				return err
			}
		case "go":
			scripts = []string{}
		default:
			scripts, err = manager.ListScripts()
			if err != nil {
				scripts = []string{}
			}
		}
		
		commands, err := manager.ListCommands()
		if err != nil {
			return err
		}
		
		if len(scripts) > 0 {
			sv.Config.Title = "Scripts & Commands"
		} else {
			sv.Config.Title = "Commands"
		}
		
		sv.Items = append(scripts, commands...)
		sv.commands = commands
		
		return nil
	})
}

func (sv *ScriptsView) CanExecute() bool {
	return len(sv.Items) > 0
}

func (sv *ScriptsView) Execute(g *gocui.Gui) error {
	if !sv.CanExecute() {
		return nil
	}
	
	selectedItem := sv.Items[sv.SelectedIndex]
	scriptsCount := len(sv.Items) - len(sv.commands)
	isCommand := sv.SelectedIndex >= scriptsCount
	
	if outputView, err := g.View("output"); err == nil {
		if isCommand {
			fmt.Fprintf(outputView, "\n=== Executing command: %s ===\n", selectedItem)
		} else {
			fmt.Fprintf(outputView, "\n=== Executing script: %s ===\n", selectedItem)
		}
		
		go func() {
			if isCommand {
				if err := sv.executeCommand(selectedItem); err != nil {
					fmt.Fprintf(outputView, "Error: %v\n", err)
				} else {
					fmt.Fprintf(outputView, "Command executed successfully\n")
				}
			} else {
				if err := helpers.WithManager("", func(manager models.PackageManager) error {
					return manager.RunScript(selectedItem, []string{})
				}); err != nil {
					fmt.Fprintf(outputView, "Error: %v\n", err)
				} else {
					fmt.Fprintf(outputView, "Script executed successfully\n")
				}
			}
		}()
	}
	
	return nil
}

func (sv *ScriptsView) executeCommand(command string) error {
	return helpers.WithManager("", func(manager models.PackageManager) error {
		switch command {
		case "init":
			return fmt.Errorf("init command should be called from CLI")
		case "build":
			return fmt.Errorf("build command should be called from CLI")
		case "test":
			return fmt.Errorf("test command should be called from CLI")
		case "clean":
			return fmt.Errorf("clean command should be called from CLI")
		case "update":
			return fmt.Errorf("update command should be called from CLI")
		default:
			if manager.Name() == "go" {
				return sv.executeGoCommand(command)
			}
			return fmt.Errorf("command not supported: %s", command)
		}
	})
}

func (sv *ScriptsView) executeGoCommand(command string) error {
	return fmt.Errorf("Go command execution not implemented yet: %s", command)
}

func (sv *ScriptsView) Refresh() error {
	sv.Items = nil
	sv.commands = nil
	return sv.loadScriptsAndCommands()
}