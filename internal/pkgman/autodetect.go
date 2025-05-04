package pkgman

import (
	"errors"
	"os"
	"path/filepath"
)

func DetectPackageManager(startDir string) (PackageManager, error) {
	lockFiles := map[string]func(string) PackageManager{
		"package-lock.json": NewNPM,
		"go.mod":            NewGoPM,
	}

	dir := startDir
	for {
		for file, constructor := range lockFiles {
			if _, err := os.Stat(filepath.Join(dir, file)); err == nil {
				return constructor(dir), nil
			}
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return nil, errors.New("no package manager detected")
}
