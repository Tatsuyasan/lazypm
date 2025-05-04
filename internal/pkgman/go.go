package pkgman

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type GoPM struct {
	Dir string
}

func NewGoPM(dir string) PackageManager {
	return &GoPM{Dir: dir}
}

func (g *GoPM) Name() string {
	return "go"
}

func (g *GoPM) Install(args []string) error {
	cmdArgs := append([]string{"mod", "tidy"}, args...)
	cmd := exec.Command("go", cmdArgs...)
	cmd.Dir = g.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (g *GoPM) RunScript(script string, args []string) error {
	return fmt.Errorf("go ne supporte pas les scripts personnalisés")
}

func (g *GoPM) ListScripts() ([]string, error) {
	// Pas de scripts personnalisés en Go
	return []string{}, nil
}

func (g *GoPM) ListDependencies() ([]string, error) {
	cmd := exec.Command("go", "list", "-m", "all")
	cmd.Dir = g.Dir
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")
	if len(lines) <= 1 {
		return []string{}, nil
	}
	// Skip first line (main module)
	return lines[1:], nil
}
