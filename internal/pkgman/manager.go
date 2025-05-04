package pkgman

type PackageManager interface {
	Name() string
	RunScript(script string, args []string) error
	ListScripts() ([]string, error)
	ListDependencies() ([]string, error)
	Install(args []string) error
}
