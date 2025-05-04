package pkgman

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
)

type NPM struct {
	Dir string
}

func NewNPM(dir string) PackageManager {
	return &NPM{Dir: dir}
}

func (n *NPM) Name() string {
	return "npm"
}

func (n *NPM) Install(args []string) error {
	cmdArgs := append([]string{"install"}, args...)
	cmd := exec.Command("npm", cmdArgs...)
	cmd.Dir = n.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (n *NPM) RunScript(script string, args []string) error {
	cmdArgs := append([]string{"run", script}, args...)
	cmd := exec.Command("npm", cmdArgs...)
	cmd.Dir = n.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (n *NPM) ListScripts() ([]string, error) {
	pkgJsonPath := filepath.Join(n.Dir, "package.json")
	data, err := os.ReadFile(pkgJsonPath)
	if err != nil {
		return nil, err
	}

	var pkg struct {
		Scripts map[string]string `json:"scripts"`
	}
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
	cmd := exec.Command("npm", "ls", "--depth=0", "--json")
	cmd.Dir = n.Dir
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var parsed struct {
		Dependencies map[string]any `json:"dependencies"`
	}
	if err := json.Unmarshal(out, &parsed); err != nil {
		return nil, err
	}

	var deps []string
	for name := range parsed.Dependencies {
		deps = append(deps, name)
	}
	return deps, nil
}
