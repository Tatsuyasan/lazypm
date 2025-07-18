package views

import (
	"fmt"

	"github.com/Tatsuyasan/lazyPm/packages/gui/components"
	"github.com/Tatsuyasan/lazyPm/packages/helpers"
	"github.com/Tatsuyasan/lazyPm/packages/models"
	"github.com/jroimartin/gocui"
)

type PackageManagerView struct {
	*components.BasePanel
	manager models.PackageManager
}

func NewPackageManagerView() *PackageManagerView {
	config := components.BasePanelConfig{
		Name:        "package_manager",
		Title:       "Package Manager",
		Highlighted: true,
		Selectable:  false,
	}
	
	return &PackageManagerView{
		BasePanel: components.NewBasePanel(config),
	}
}

func (pmv *PackageManagerView) Render(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}
	
	v.Clear()
	v.Title = pmv.Config.Title
	
	if pmv.manager == nil {
		if err := helpers.WithManager("", func(manager models.PackageManager) error {
			pmv.manager = manager
			return nil
		}); err != nil {
			fmt.Fprintf(v, "Error detecting package manager: %v", err)
			return nil
		}
	}
	
	if pmv.manager != nil {
		fmt.Fprintf(v, "Detected: %s\n", pmv.manager.Name())
		
		switch pmv.manager.Name() {
		case "npm":
			fmt.Fprintf(v, "Type: Node.js Package Manager\n")
			fmt.Fprintf(v, "Config: package.json\n")
		case "go":
			fmt.Fprintf(v, "Type: Go Modules\n")
			fmt.Fprintf(v, "Config: go.mod\n")
		default:
			fmt.Fprintf(v, "Type: Unknown\n")
		}
	} else {
		fmt.Fprintf(v, "No package manager detected")
	}
	
	return nil
}

func (pmv *PackageManagerView) HandleKey(g *gocui.Gui, v *gocui.View, key interface{}) error {
	return nil
}

func (pmv *PackageManagerView) GetManager() models.PackageManager {
	return pmv.manager
}