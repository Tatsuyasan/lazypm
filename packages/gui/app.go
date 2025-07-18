package gui

import (
	"log"

	"github.com/Tatsuyasan/lazyPm/packages/gui/components"
	"github.com/Tatsuyasan/lazyPm/packages/gui/config"
	"github.com/Tatsuyasan/lazyPm/packages/gui/handlers"
	"github.com/Tatsuyasan/lazyPm/packages/gui/views"
	"github.com/jroimartin/gocui"
)

type LazyPmGUI struct {
	gui           *gocui.Gui
	layoutManager *components.LayoutManager
	keyBindings   *config.KeyBindings
	keyHandler    *handlers.KeyBindingHandler
}

func RunGUI() error {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	lazyPmGUI := NewLazyPmGUI(g)

	if err := lazyPmGUI.Setup(); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

	return nil
}

func NewLazyPmGUI(g *gocui.Gui) *LazyPmGUI {
	layoutManager := components.NewLayoutManager()
	keyBindings := config.GetDefaultKeyBindings()
	keyHandler := handlers.NewKeyBindingHandler(layoutManager, keyBindings)

	return &LazyPmGUI{
		gui:           g,
		layoutManager: layoutManager,
		keyBindings:   keyBindings,
		keyHandler:    keyHandler,
	}
}

func (lpg *LazyPmGUI) Setup() error {
	lpg.gui.Highlight = true
	lpg.gui.SelFgColor = gocui.ColorGreen
	lpg.gui.Cursor = true

	lpg.gui.FgColor = gocui.ColorWhite

	lpg.createPanels()

	lpg.gui.SetManagerFunc(lpg.layoutManager.Layout)

	if err := lpg.keyHandler.SetupKeyBindings(lpg.gui); err != nil {
		return err
	}

	return nil
}

func (lpg *LazyPmGUI) createPanels() {
	packageManagerView := views.NewPackageManagerView()
	packagesView := views.NewPackagesView()
	dependenciesView := views.NewDependenciesView()
	scriptsView := views.NewScriptsView()

	lpg.layoutManager.AddPanel(packageManagerView)
	lpg.layoutManager.AddPanel(packagesView)
	lpg.layoutManager.AddPanel(dependenciesView)
	lpg.layoutManager.AddPanel(scriptsView)
}
