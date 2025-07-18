package helpers

import (
	"fmt"
	"os"

	"github.com/Tatsuyasan/lazyPm/packages/models"
	"github.com/Tatsuyasan/lazyPm/packages/pkgman"
)

func WithManager(pmFlag string, cb func(models.PackageManager) error) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	manager, err := getPackageManagerByFlagManager(pmFlag, dir)
	if err != nil {
		return err
	}

	return cb(manager)
}

func getPackageManagerByFlagManager(pmFlag, dir string) (models.PackageManager, error) {
	if pmFlag != "" {
		switch pmFlag {
		case "npm":
			return pkgman.NewNPM(dir), nil
		case "go":
			return pkgman.NewGoPM(dir), nil
		default:
			return nil, fmt.Errorf("unknown package manager: %s", pmFlag)
		}
	}

	manager, err := DetectPackageManager(dir)
	if err != nil {
		return nil, err
	}

	return manager, nil
}
