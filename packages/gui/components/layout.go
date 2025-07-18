package components

import (
	"github.com/jroimartin/gocui"
)

type LayoutManager struct {
	panels       []Panel
	currentPanel int
	maxX, maxY   int
}

func NewLayoutManager() *LayoutManager {
	return &LayoutManager{
		panels:       make([]Panel, 0),
		currentPanel: 0,
	}
}

func (lm *LayoutManager) AddPanel(panel Panel) {
	lm.panels = append(lm.panels, panel)
}

func (lm *LayoutManager) GetCurrentPanel() Panel {
	if lm.currentPanel >= 0 && lm.currentPanel < len(lm.panels) {
		return lm.panels[lm.currentPanel]
	}
	return nil
}

func (lm *LayoutManager) NextPanel(g *gocui.Gui) error {
	if len(lm.panels) == 0 {
		return nil
	}

	if current := lm.GetCurrentPanel(); current != nil {
		if v, err := g.View(current.GetName()); err == nil {
			current.OnBlur(g, v)
		}
	}

	lm.currentPanel = (lm.currentPanel + 1) % len(lm.panels)

	if current := lm.GetCurrentPanel(); current != nil {
		if _, err := g.SetCurrentView(current.GetName()); err != nil {
			return err
		}
		if v, err := g.View(current.GetName()); err == nil {
			current.OnFocus(g, v)
		}
	}

	return nil
}

func (lm *LayoutManager) PrevPanel(g *gocui.Gui) error {
	if len(lm.panels) == 0 {
		return nil
	}

	if current := lm.GetCurrentPanel(); current != nil {
		if v, err := g.View(current.GetName()); err == nil {
			current.OnBlur(g, v)
		}
	}

	lm.currentPanel--
	if lm.currentPanel < 0 {
		lm.currentPanel = len(lm.panels) - 1
	}

	if current := lm.GetCurrentPanel(); current != nil {
		if _, err := g.SetCurrentView(current.GetName()); err != nil {
			return err
		}
		if v, err := g.View(current.GetName()); err == nil {
			current.OnFocus(g, v)
		}
	}

	return nil
}

func (lm *LayoutManager) Layout(g *gocui.Gui) error {
	lm.maxX, lm.maxY = g.Size()

	leftWidth := lm.maxX / 2
	
	packageManagerHeight := 5
	remainingHeight := lm.maxY - packageManagerHeight
	panelHeight := remainingHeight / 3

	leftPanels := []struct {
		name  string
		title string
		y0    int
		y1    int
	}{
		{"package_manager", "Package Manager", 0, packageManagerHeight - 1},
		{"packages", "Packages", packageManagerHeight, packageManagerHeight + panelHeight - 1},
		{"dependencies", "Dependencies", packageManagerHeight + panelHeight, packageManagerHeight + panelHeight*2 - 1},
		{"scripts", "Scripts", packageManagerHeight + panelHeight*2, lm.maxY - 1},
	}

	for _, panel := range leftPanels {
		if v, err := g.SetView(panel.name, 0, panel.y0, leftWidth-1, panel.y1); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			v.Title = panel.title
			if panel.name == "package_manager" {
				v.Highlight = false
				v.FgColor = gocui.ColorWhite
			} else {
				v.Highlight = true
				v.SelFgColor = gocui.ColorGreen
				v.FgColor = gocui.ColorWhite
			}
			v.Frame = true
		}
	}

	rightHeight := lm.maxY / 2

	if v, err := g.SetView("output", leftWidth, 0, lm.maxX-1, rightHeight-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Output"
		v.Wrap = true
		v.Autoscroll = true
	}

	if v, err := g.SetView("details", leftWidth, rightHeight, lm.maxX-1, lm.maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Details"
		v.Wrap = true
	}

	for _, panel := range lm.panels {
		if v, err := g.View(panel.GetName()); err == nil {
			if err := panel.Render(g, v); err != nil {
				return err
			}
		}
	}

	if g.CurrentView() == nil && len(lm.panels) > 0 {
		if current := lm.GetCurrentPanel(); current != nil {
			g.SetCurrentView(current.GetName())
			if v, err := g.View(current.GetName()); err == nil {
				current.OnFocus(g, v)
			}
		}
	}

	return nil
}
