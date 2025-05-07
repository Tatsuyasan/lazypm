package helpers

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Tatsuyasan/lazyPm/internal/models"
	"github.com/Tatsuyasan/lazyPm/internal/pkgman"
)

func DetectPackageManager(startDir string) (models.PackageManager, error) {
	dir := startDir
	lockFiles := map[string]func(string) models.PackageManager{
		"package-lock.json": pkgman.NewNPM,
		"go.mod":            pkgman.NewGoPM,
	}

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
