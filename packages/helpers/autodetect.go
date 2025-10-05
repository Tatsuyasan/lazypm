// Package helpers provides utility functions for detecting and
// retrieving supported package managers in lazyPm.
package helpers

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Tatsuyasan/lazyPm/packages/context"
	"github.com/Tatsuyasan/lazyPm/packages/pkgman"
)

type PackageManagerName string

const (
	Npm  PackageManagerName = "npm"
	Yarn PackageManagerName = "yarn"
	Go   PackageManagerName = "go"
)

func detectPackageManager() PackageManagerName {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	lockFiles := map[string]PackageManagerName{
		"package-lock.json": Npm,
		"yarn.lock":         Yarn,
	}

	for lockFile, managerName := range lockFiles {
		path := filepath.Join(dir, lockFile)
		if _, err := os.Stat(path); err == nil {
			return managerName
		}
	}

	return ""
}

func GetPackageManager(name PackageManagerName) (context.PackageManager, error) {
	packageManagersMap := map[PackageManagerName]func() context.PackageManager{
		Npm: func() context.PackageManager { return pkgman.NewNpm() },
		// Yarn: func() models.PackageManager { return &context.Yarn{} },
	}

	pmName := PackageManagerName("")
	if name != "" {
		pmName = name
	} else {
		pmName = detectPackageManager()
	}

	if pmName == "" {
		return nil, errors.New("no package manager detected")
	}

	getPm, ok := packageManagersMap[pmName]
	if !ok {
		return nil, errors.New("package manager not supported")
	}

	return getPm(), nil
}
