package handlers

import (
	"github.com/Tatsuyasan/lazyPm/packages/gui/components"
	"github.com/Tatsuyasan/lazyPm/packages/gui/config"
	"github.com/jroimartin/gocui"
)

type KeyBindingHandler struct {
	layoutManager *components.LayoutManager
	keyBindings   *config.KeyBindings
}

func NewKeyBindingHandler(layoutManager *components.LayoutManager, keyBindings *config.KeyBindings) *KeyBindingHandler {
	return &KeyBindingHandler{
		layoutManager: layoutManager,
		keyBindings:   keyBindings,
	}
}

func (kh *KeyBindingHandler) SetupKeyBindings(g *gocui.Gui) error {
	for action, binding := range kh.keyBindings.Global {
		switch action {
		case config.ActionQuit:
			if err := g.SetKeybinding("", binding.Key, binding.Mod, kh.handleQuit); err != nil {
				return err
			}
		case config.ActionForceQuit:
			if err := g.SetKeybinding("", binding.Key, binding.Mod, kh.handleForceQuit); err != nil {
				return err
			}
		case config.ActionNextPanel:
			if err := g.SetKeybinding("", binding.Key, binding.Mod, kh.handleNextPanel); err != nil {
				return err
			}
		case config.ActionPrevPanel:
			if err := g.SetKeybinding("", binding.Key, binding.Mod, kh.handlePrevPanel); err != nil {
				return err
			}
		case config.ActionMoveUp:
			if err := g.SetKeybinding("", binding.Key, binding.Mod, kh.handleMoveUp); err != nil {
				return err
			}
		case config.ActionMoveDown:
			if err := g.SetKeybinding("", binding.Key, binding.Mod, kh.handleMoveDown); err != nil {
				return err
			}
		case config.ActionExecute:
			if err := g.SetKeybinding("", binding.Key, binding.Mod, kh.handleExecute); err != nil {
				return err
			}
		case config.ActionRefresh:
			if err := g.SetKeybinding("", binding.Key, binding.Mod, kh.handleRefresh); err != nil {
				return err
			}
		}
	}
	
	if err := g.SetKeybinding("", 'l', gocui.ModNone, kh.handleNextPanel); err != nil {
		return err
	}
	
	return nil
}

func (kh *KeyBindingHandler) handleQuit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (kh *KeyBindingHandler) handleForceQuit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (kh *KeyBindingHandler) handleNextPanel(g *gocui.Gui, v *gocui.View) error {
	return kh.layoutManager.NextPanel(g)
}

func (kh *KeyBindingHandler) handlePrevPanel(g *gocui.Gui, v *gocui.View) error {
	return kh.layoutManager.PrevPanel(g)
}

func (kh *KeyBindingHandler) handleMoveUp(g *gocui.Gui, v *gocui.View) error {
	if panel := kh.layoutManager.GetCurrentPanel(); panel != nil {
		if err := panel.HandleKey(g, v, 'k'); err != nil {
			return err
		}
		if currentView, err := g.View(panel.GetName()); err == nil {
			return panel.Render(g, currentView)
		}
	}
	return nil
}

func (kh *KeyBindingHandler) handleMoveDown(g *gocui.Gui, v *gocui.View) error {
	if panel := kh.layoutManager.GetCurrentPanel(); panel != nil {
		if err := panel.HandleKey(g, v, 'j'); err != nil {
			return err
		}
		if currentView, err := g.View(panel.GetName()); err == nil {
			return panel.Render(g, currentView)
		}
	}
	return nil
}

func (kh *KeyBindingHandler) handleExecute(g *gocui.Gui, v *gocui.View) error {
	if panel := kh.layoutManager.GetCurrentPanel(); panel != nil {
		return panel.HandleKey(g, v, gocui.KeyEnter)
	}
	return nil
}

func (kh *KeyBindingHandler) handleRefresh(g *gocui.Gui, v *gocui.View) error {
	if panel := kh.layoutManager.GetCurrentPanel(); panel != nil {
		if refreshable, ok := panel.(interface{ Refresh() error }); ok {
			return refreshable.Refresh()
		}
	}
	return nil
}