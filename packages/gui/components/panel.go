package components

import (
	"github.com/jroimartin/gocui"
)

type Panel interface {
	GetName() string

	Render(g *gocui.Gui, v *gocui.View) error

	OnFocus(g *gocui.Gui, v *gocui.View) error

	OnBlur(g *gocui.Gui, v *gocui.View) error

	HandleKey(g *gocui.Gui, v *gocui.View, key interface{}) error

	GetItems() []string

	GetSelectedIndex() int

	SetSelectedIndex(index int)

	CanExecute() bool

	Execute(g *gocui.Gui) error
}

type BasePanelConfig struct {
	Name        string
	Title       string
	Highlighted bool
	Selectable  bool
}

type BasePanel struct {
	Config        BasePanelConfig
	SelectedIndex int
	Items         []string
	ScrollOffset  int
}

func NewBasePanel(config BasePanelConfig) *BasePanel {
	return &BasePanel{
		Config:        config,
		SelectedIndex: 0,
		Items:         make([]string, 0),
		ScrollOffset:  0,
	}
}

func (bp *BasePanel) GetName() string {
	return bp.Config.Name
}

func (bp *BasePanel) GetItems() []string {
	return bp.Items
}

func (bp *BasePanel) GetSelectedIndex() int {
	return bp.SelectedIndex
}

func (bp *BasePanel) SetSelectedIndex(index int) {
	if index >= 0 && index < len(bp.Items) {
		bp.SelectedIndex = index
	}
}

func (bp *BasePanel) MoveUp() {
	if bp.SelectedIndex > 0 {
		bp.SelectedIndex--
	}
}

func (bp *BasePanel) MoveDown() {
	if bp.SelectedIndex < len(bp.Items)-1 {
		bp.SelectedIndex++
	}
}

func (bp *BasePanel) MoveUpWithScroll(viewHeight int) {
	if bp.SelectedIndex > 0 {
		bp.SelectedIndex--
		bp.AdjustScrollOffset(viewHeight)
	}
}

func (bp *BasePanel) MoveDownWithScroll(viewHeight int) {
	if bp.SelectedIndex < len(bp.Items)-1 {
		bp.SelectedIndex++
		bp.AdjustScrollOffset(viewHeight)
	}
}

func (bp *BasePanel) OnFocus(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		v.Highlight = true
		v.FgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorWhite
		v.SelBgColor = gocui.ColorBlue
	}
	return nil
}

func (bp *BasePanel) OnBlur(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		v.Highlight = false
		v.FgColor = gocui.ColorWhite
		v.SelFgColor = gocui.ColorDefault
		v.SelBgColor = gocui.ColorDefault
	}
	return nil
}

func (bp *BasePanel) CanExecute() bool {
	return false
}

func (bp *BasePanel) Execute(g *gocui.Gui) error {
	return nil
}

func (bp *BasePanel) Render(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}
	
	if g.CurrentView() == v {
		v.SelFgColor = gocui.ColorWhite
		v.SelBgColor = gocui.ColorBlue
	} else {
		v.SelFgColor = gocui.ColorDefault
		v.SelBgColor = gocui.ColorDefault
	}
	
	return nil
}

func (bp *BasePanel) GetVisibleItems(viewHeight int) []string {
	if len(bp.Items) == 0 {
		return bp.Items
	}

	start := bp.ScrollOffset
	end := start + viewHeight

	if start < 0 {
		start = 0
	}
	if end > len(bp.Items) {
		end = len(bp.Items)
	}

	return bp.Items[start:end]
}

func (bp *BasePanel) GetVisibleSelectedIndex() int {
	return bp.SelectedIndex - bp.ScrollOffset
}

func (bp *BasePanel) AdjustScrollOffset(viewHeight int) {
	if len(bp.Items) == 0 || viewHeight <= 0 {
		return
	}

	if bp.SelectedIndex < bp.ScrollOffset {
		bp.ScrollOffset = bp.SelectedIndex
	}

	if bp.SelectedIndex >= bp.ScrollOffset+viewHeight {
		bp.ScrollOffset = bp.SelectedIndex - viewHeight + 1
	}

	if bp.ScrollOffset < 0 {
		bp.ScrollOffset = 0
	}

	maxOffset := len(bp.Items) - viewHeight
	if maxOffset < 0 {
		maxOffset = 0
	}
	if bp.ScrollOffset > maxOffset {
		bp.ScrollOffset = maxOffset
	}
}
