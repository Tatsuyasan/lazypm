// Package context provides shared application context and state for lazyPm.
package context

import (
	"sync"
)

type AppContext struct {
	Manager PackageManager
}

var (
	appCtx *AppContext
	once   sync.Once
)

func GetContext() *AppContext {
	once.Do(func() {
		appCtx = &AppContext{}
	})
	return appCtx
}
