package config

import (
	"github.com/jroimartin/gocui"
)

type KeyAction string

const (
	ActionMoveUp    KeyAction = "move_up"
	ActionMoveDown  KeyAction = "move_down"
	ActionNextPanel KeyAction = "next_panel"
	ActionPrevPanel KeyAction = "prev_panel"

	ActionExecute   KeyAction = "execute"
	ActionQuit      KeyAction = "quit"
	ActionForceQuit KeyAction = "force_quit"

	ActionRefresh    KeyAction = "refresh"
	ActionToggleHelp KeyAction = "toggle_help"
)

type KeyBinding struct {
	Key    interface{}
	Mod    gocui.Modifier
	Action KeyAction
}

type KeyBindings struct {
	Global map[KeyAction]KeyBinding
	Panel  map[string]map[KeyAction]KeyBinding
}

func GetDefaultKeyBindings() *KeyBindings {
	return &KeyBindings{
		Global: map[KeyAction]KeyBinding{
			ActionMoveUp:   {Key: 'k', Mod: gocui.ModNone, Action: ActionMoveUp},
			ActionMoveDown: {Key: 'j', Mod: gocui.ModNone, Action: ActionMoveDown},

			ActionNextPanel: {Key: gocui.KeyTab, Mod: gocui.ModNone, Action: ActionNextPanel},
			ActionPrevPanel: {Key: 'h', Mod: gocui.ModNone, Action: ActionPrevPanel},

			ActionExecute:   {Key: gocui.KeyEnter, Mod: gocui.ModNone, Action: ActionExecute},
			ActionQuit:      {Key: 'q', Mod: gocui.ModNone, Action: ActionQuit},
			ActionForceQuit: {Key: gocui.KeyCtrlC, Mod: gocui.ModNone, Action: ActionForceQuit},

					ActionRefresh:    {Key: gocui.KeyF5, Mod: gocui.ModNone, Action: ActionRefresh},
			ActionToggleHelp: {Key: '?', Mod: gocui.ModNone, Action: ActionToggleHelp},
		},
		Panel: make(map[string]map[KeyAction]KeyBinding),
	}
}

func (kb *KeyBindings) GetKeyBinding(action KeyAction) (KeyBinding, bool) {
	binding, exists := kb.Global[action]
	return binding, exists
}

func (kb *KeyBindings) GetPanelKeyBinding(panelName string, action KeyAction) (KeyBinding, bool) {
	if panelBindings, exists := kb.Panel[panelName]; exists {
		if binding, exists := panelBindings[action]; exists {
			return binding, true
		}
	}
	return kb.GetKeyBinding(action)
}

func (kb *KeyBindings) AddPanelKeyBinding(panelName string, action KeyAction, binding KeyBinding) {
	if kb.Panel[panelName] == nil {
		kb.Panel[panelName] = make(map[KeyAction]KeyBinding)
	}
	kb.Panel[panelName][action] = binding
}
