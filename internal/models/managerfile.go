package models

type PackageManagerFile struct {
	Name             string            `json:"name"`
	Version          string            `json:"version"`
	Description      string            `json:"description,omitempty"`
	Main             string            `json:"main,omitempty"`
	Scripts          map[string]string `json:"scripts,omitempty"`
	Dependencies     map[string]string `json:"dependencies,omitempty"`
	DevDependencies  map[string]string `json:"devDependencies,omitempty"`
	PeerDependencies map[string]string `json:"peerDependencies,omitempty"`
	Keywords         []string          `json:"keywords,omitempty"`
	Author           string            `json:"author,omitempty"`
	License          string            `json:"license,omitempty"`
	Private          bool              `json:"private,omitempty"`
	Engines          map[string]string `json:"engines,omitempty"`
}
