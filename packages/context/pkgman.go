package context

type ScriptInfo struct {
	Name        string
	Description string
}

type PackageManager interface {
	Name() string
	RunScript(script string, args []string) error
	ListScripts() ([]ScriptInfo, error)
	ListDependencies() ([]string, error)
	Install(args []string) error
	ListCommands() ([]string, error)
}
