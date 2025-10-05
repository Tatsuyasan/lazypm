// Package pkgman defines the interface and implementations for package managers.
package pkgman

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Tatsuyasan/lazyPm/packages/context"
	"github.com/Tatsuyasan/lazyPm/packages/models"
)

type Npm struct{}

func NewNpm() *Npm {
	return &Npm{}
}

// Install implements models.PackageManager.
func (n Npm) Install(args []string) error {
	ctx := context.GetContext()

	cmdArgs := append([]string{"install"}, args...)
	cmd := exec.Command(ctx.Manager.Name(), cmdArgs...)
	cmd.Dir = "."
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ListCommands implements models.PackageManager.
func (n Npm) ListCommands() ([]string, error) {
	panic("unimplemented")
}

// ListDependencies implements models.PackageManager.
func (n Npm) ListDependencies() ([]string, error) {
	panic("unimplemented")
}

// ListScripts implements models.PackageManager.
func (n Npm) ListScripts() ([]context.ScriptInfo, error) {
	path := filepath.Join(".", "package.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var pkg models.PackageManagerFile

	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, err
	}

	var scripts []context.ScriptInfo
	for name, desc := range pkg.Scripts {
		scripts = append(scripts, context.ScriptInfo{Name: name, Description: desc})
	}

	return scripts, nil
}

// Name implements models.PackageManager.
func (n Npm) Name() string {
	return "npm"
}

// RunScript implements models.PackageManager.
func (n Npm) RunScript(script string, args []string) error {
	panic("unimplemented")
}
