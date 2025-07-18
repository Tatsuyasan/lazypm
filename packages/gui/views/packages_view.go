package views

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Tatsuyasan/lazyPm/packages/gui/components"
	"github.com/Tatsuyasan/lazyPm/packages/helpers"
	"github.com/Tatsuyasan/lazyPm/packages/models"
	"github.com/jroimartin/gocui"
)

type PackagesView struct {
	*components.BasePanel
	manager     models.PackageManager
	packageInfo map[string]string
}

func NewPackagesView() *PackagesView {
	config := components.BasePanelConfig{
		Name:        "packages",
		Title:       "Packages",
		Highlighted: true,
		Selectable:  true,
	}
	
	return &PackagesView{
		BasePanel:   components.NewBasePanel(config),
		packageInfo: make(map[string]string),
	}
}

func (pv *PackagesView) Render(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}
	
	pv.BasePanel.Render(g, v)
	
	v.Clear()
	v.Title = pv.Config.Title
	
	if len(pv.Items) == 0 {
		if err := pv.loadPackages(); err != nil {
			fmt.Fprintf(v, "Error loading packages: %v", err)
			return nil
		}
	}
	
	if len(pv.Items) == 0 {
		fmt.Fprintf(v, "No packages found")
		return nil
	}
	
	_, viewHeight := v.Size()
	viewHeight -= 2
	effectiveHeight := viewHeight / 2
	
	pv.AdjustScrollOffset(effectiveHeight)
	
	visibleItems := pv.GetVisibleItems(effectiveHeight)
	visibleSelectedIndex := pv.GetVisibleSelectedIndex()
	
	viewWidth, _ := v.Size()
	for _, pkg := range visibleItems {
		if path, exists := pv.packageInfo[pkg]; exists {
			pkgPadding := viewWidth - len(pkg)
			pkgContent := pkg
			if pkgPadding > 0 {
				pkgContent += fmt.Sprintf("%*s", pkgPadding, "")
			}
			fmt.Fprintf(v, "%s\n", pkgContent)
			
			pathContent := "    " + path
			pathPadding := viewWidth - len(pathContent)
			if pathPadding > 0 {
				pathContent += fmt.Sprintf("%*s", pathPadding, "")
			}
			fmt.Fprintf(v, "%s\n", pathContent)
		} else {
			pkgPadding := viewWidth - len(pkg)
			pkgContent := pkg
			if pkgPadding > 0 {
				pkgContent += fmt.Sprintf("%*s", pkgPadding, "")
			}
			fmt.Fprintf(v, "%s\n", pkgContent)
		}
	}
	
	if len(visibleItems) > 0 && visibleSelectedIndex >= 0 && visibleSelectedIndex < len(visibleItems) {
		cursorY := visibleSelectedIndex * 2
		v.SetCursor(0, cursorY)
	}
	
	return nil
}

func (pv *PackagesView) HandleKey(g *gocui.Gui, v *gocui.View, key interface{}) error {
	_, viewHeight := v.Size()
	viewHeight -= 2
	effectiveHeight := viewHeight / 2
	
	switch key {
	case 'j':
		pv.MoveDownWithScroll(effectiveHeight)
	case 'k':
		pv.MoveUpWithScroll(effectiveHeight)
	case gocui.KeyEnter:
		return pv.showDetails(g)
	}
	return nil
}

func (pv *PackagesView) loadPackages() error {
	return helpers.WithManager("", func(manager models.PackageManager) error {
		pv.manager = manager
		
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		
		switch manager.Name() {
		case "npm":
			return pv.loadNpmPackages(cwd)
		case "go":
			return pv.loadGoPackages(cwd)
		default:
			pv.Items = []string{"Current Project"}
			pv.packageInfo["Current Project"] = cwd
		}
		
		return nil
	})
}

func (pv *PackagesView) loadNpmPackages(cwd string) error {
	packageJsonPath := filepath.Join(cwd, "package.json")
	if _, err := os.Stat(packageJsonPath); err == nil {
		projectName := filepath.Base(cwd)
		pv.Items = []string{projectName}
		pv.packageInfo[projectName] = cwd
	}
	
	return nil
}

func (pv *PackagesView) loadGoPackages(cwd string) error {
	goModPath := filepath.Join(cwd, "go.mod")
	if _, err := os.Stat(goModPath); err == nil {
		projectName := filepath.Base(cwd)
		pv.Items = []string{projectName}
		pv.packageInfo[projectName] = cwd
	}
	
	return nil
}

func (pv *PackagesView) showDetails(g *gocui.Gui) error {
	if len(pv.Items) == 0 {
		return nil
	}
	
	selectedPkg := pv.Items[pv.SelectedIndex]
	
	if detailsView, err := g.View("details"); err == nil {
		detailsView.Clear()
		fmt.Fprintf(detailsView, "Package: %s\n", selectedPkg)
		
		if path, exists := pv.packageInfo[selectedPkg]; exists {
			fmt.Fprintf(detailsView, "Path: %s\n", path)
			
			switch pv.manager.Name() {
			case "npm":
				pv.showNpmPackageDetails(detailsView, path)
			case "go":
				pv.showGoPackageDetails(detailsView, path)
			}
		}
	}
	
	return nil
}

func (pv *PackagesView) showNpmPackageDetails(v *gocui.View, path string) {
	fmt.Fprintf(v, "\nNPM Package Details:\n")
	
	packageJsonPath := filepath.Join(path, "package.json")
	if _, err := os.Stat(packageJsonPath); err == nil {
		fmt.Fprintf(v, "- package.json: ✓\n")
	} else {
		fmt.Fprintf(v, "- package.json: ✗\n")
	}
	
	nodeModulesPath := filepath.Join(path, "node_modules")
	if _, err := os.Stat(nodeModulesPath); err == nil {
		fmt.Fprintf(v, "- node_modules: ✓\n")
	} else {
		fmt.Fprintf(v, "- node_modules: ✗\n")
	}
}

func (pv *PackagesView) showGoPackageDetails(v *gocui.View, path string) {
	fmt.Fprintf(v, "\nGo Package Details:\n")
	
	goModPath := filepath.Join(path, "go.mod")
	if _, err := os.Stat(goModPath); err == nil {
		fmt.Fprintf(v, "- go.mod: ✓\n")
	} else {
		fmt.Fprintf(v, "- go.mod: ✗\n")
	}
	
	goSumPath := filepath.Join(path, "go.sum")
	if _, err := os.Stat(goSumPath); err == nil {
		fmt.Fprintf(v, "- go.sum: ✓\n")
	} else {
		fmt.Fprintf(v, "- go.sum: ✗\n")
	}
}

func (pv *PackagesView) Refresh() error {
	pv.Items = nil
	pv.packageInfo = make(map[string]string)
	return pv.loadPackages()
}