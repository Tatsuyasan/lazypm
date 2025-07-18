package views

import (
	"fmt"

	"github.com/Tatsuyasan/lazyPm/packages/gui/components"
	"github.com/Tatsuyasan/lazyPm/packages/helpers"
	"github.com/Tatsuyasan/lazyPm/packages/models"
	"github.com/jroimartin/gocui"
)

type DependenciesView struct {
	*components.BasePanel
	manager models.PackageManager
}

func NewDependenciesView() *DependenciesView {
	config := components.BasePanelConfig{
		Name:        "dependencies",
		Title:       "Dependencies",
		Highlighted: true,
		Selectable:  true,
	}
	
	return &DependenciesView{
		BasePanel: components.NewBasePanel(config),
	}
}

func (dv *DependenciesView) Render(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}
	
	dv.BasePanel.Render(g, v)
	
	v.Clear()
	v.Title = dv.Config.Title
	
	if len(dv.Items) == 0 {
		if err := dv.loadDependencies(); err != nil {
			fmt.Fprintf(v, "Error loading dependencies: %v", err)
			return nil
		}
	}
	
	if len(dv.Items) == 0 {
		fmt.Fprintf(v, "No dependencies found")
		return nil
	}
	
	_, viewHeight := v.Size()
	viewHeight -= 2
	
	dv.AdjustScrollOffset(viewHeight)
	
	visibleItems := dv.GetVisibleItems(viewHeight)
	visibleSelectedIndex := dv.GetVisibleSelectedIndex()
	
	viewWidth, _ := v.Size()
	for _, dep := range visibleItems {
		padding := viewWidth - len(dep)
		content := dep
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

func (dv *DependenciesView) HandleKey(g *gocui.Gui, v *gocui.View, key interface{}) error {
	_, viewHeight := v.Size()
	viewHeight -= 2
	
	switch key {
	case 'j':
		dv.MoveDownWithScroll(viewHeight)
	case 'k':
		dv.MoveUpWithScroll(viewHeight)
	case gocui.KeyEnter:
		return dv.showDetails(g)
	}
	return nil
}

func (dv *DependenciesView) loadDependencies() error {
	return helpers.WithManager("", func(manager models.PackageManager) error {
		dv.manager = manager
		deps, err := manager.ListDependencies()
		if err != nil {
			return err
		}
		dv.Items = deps
		return nil
	})
}

func (dv *DependenciesView) showDetails(g *gocui.Gui) error {
	if len(dv.Items) == 0 {
		return nil
	}
	
	selectedDep := dv.Items[dv.SelectedIndex]
	
	if detailsView, err := g.View("details"); err == nil {
		detailsView.Clear()
		fmt.Fprintf(detailsView, "Dependency: %s\n", selectedDep)
		fmt.Fprintf(detailsView, "\nDetails:\n")
		fmt.Fprintf(detailsView, "- Package manager: %s\n", dv.manager.Name())
		fmt.Fprintf(detailsView, "- Status: Installed\n")
		
	}
	
	return nil
}

func (dv *DependenciesView) Refresh() error {
	dv.Items = nil
	return dv.loadDependencies()
}