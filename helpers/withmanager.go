package helpers

import (
	"fmt"
	"os"

	"github.com/Tatsuyasan/lazyPm/internal/pkgman"
)

func WithManager(pmFlag string, cb func(pkgman.PackageManager) error) error {
	dir, _ := os.Getwd()

	manager, err := getPackageManagerByFlagManager(pmFlag, dir)
	if err != nil {
		return err
	}

	return cb(manager)
}

func getPackageManagerByFlagManager(pmFlag, dir string) (pkgman.PackageManager, error) {
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

	manager, err := pkgman.DetectPackageManager(dir)
	if err != nil {
		return nil, err
	}

	return manager, nil
}
