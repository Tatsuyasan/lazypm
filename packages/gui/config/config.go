package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	KeyBindings *KeyBindings
	Theme       *Theme
	Behavior    *Behavior
}

type Theme struct {
	HighlightColor  string
	SelectedColor   string
	BackgroundColor string
	TextColor       string
	BorderColor     string
	ErrorColor      string
	SuccessColor    string
}

type Behavior struct {
	AutoRefresh      bool
	RefreshInterval  int
	ShowHelp         bool
	ConfirmExecution bool
	ScrollOffset     int
}

func GetDefaultConfig() *Config {
	return &Config{
		KeyBindings: GetDefaultKeyBindings(),
		Theme:       GetDefaultTheme(),
		Behavior:    GetDefaultBehavior(),
	}
}

func GetDefaultTheme() *Theme {
	return &Theme{
		HighlightColor:  "blue",
		SelectedColor:   "white",
		BackgroundColor: "black",
		TextColor:       "white",
		BorderColor:     "gray",
		ErrorColor:      "red",
		SuccessColor:    "green",
	}
}

func GetDefaultBehavior() *Behavior {
	return &Behavior{
		AutoRefresh:      false,
		RefreshInterval:  30,
		ShowHelp:         true,
		ConfirmExecution: false,
		ScrollOffset:     3,
	}
}

func GetConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ".lazypm"
	}
	return filepath.Join(home, ".config", "lazypm")
}

func GetConfigFile() string {
	return filepath.Join(GetConfigPath(), "config.yaml")
}

func GetKeybindingsFile() string {
	return filepath.Join(GetConfigPath(), "keybindings.yaml")
}
