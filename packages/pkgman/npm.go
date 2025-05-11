package pkgman

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Tatsuyasan/lazyPm/packages/models"
)

const (
	packageManagerFile = "package.json"
	packageManagerCmd  = "npm"
)

type NPM struct {
	Dir string
}

func NewNPM(dir string) models.PackageManager {
	return &NPM{Dir: dir}
}

func (n *NPM) Name() string {
	return packageManagerCmd
}

func (n *NPM) Install(args []string) error {
	cmdArgs := append([]string{"install"}, args...)
	cmd := exec.Command(packageManagerCmd, cmdArgs...)
	cmd.Dir = n.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (n *NPM) RunScript(script string, args []string) error {
	cmdArgs := append([]string{"run", script}, args...)
	cmd := exec.Command(packageManagerCmd, cmdArgs...)
	cmd.Dir = n.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (n *NPM) ListScripts() ([]string, error) {
	pkgJsonPath := filepath.Join(n.Dir, packageManagerFile)
	data, err := os.ReadFile(pkgJsonPath)
	if err != nil {
		return nil, err
	}

	var pkg models.PackageManagerFile

	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, err
	}

	var scripts []string
	for name := range pkg.Scripts {
		scripts = append(scripts, name)
	}
	return scripts, nil
}

func (n *NPM) ListDependencies() ([]string, error) {
	data, err := n.readFile(packageManagerFile)
	if err != nil {
		return nil, err
	}

	var parsed models.PackageManagerFile

	if err := json.Unmarshal(data, &parsed); err != nil {
		return nil, err
	}

	var dependencies []string
	for name := range parsed.Dependencies {
		dependencies = append(dependencies, name)
	}
	return dependencies, nil
}

func (n *NPM) readFile(filename string) ([]byte, error) {
	packageJSONPath := filepath.Join(n.Dir, filename)
	return os.ReadFile(packageJSONPath)
}
