package gui

import (
	"fmt"
	"log"
	"os"

	"github.com/jroimartin/gocui"
)

var (
	views            = []string{"left", "right"}
	currentViewIndex int
	firstLaunch      bool
)

func Init() {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}
	log.SetOutput(logFile)
}

func RunGUI() error {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
	return nil
}

func layout(g *gocui.Gui) error {
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen
	maxX, maxY := g.Size()

	// Left view
	if v, err := g.SetView("left", 0, 0, maxX/2-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Left View"
		v.Wrap = true
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlue
		fmt.Fprintln(v, "Line 1 in Left View")
		fmt.Fprintln(v, "Line 2 in Left View")
		fmt.Fprintln(v, "Line 3 in Left View")
	}

	// Right view
	if v, err := g.SetView("right", maxX/2, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Right View"
		v.Wrap = true
		v.Highlight = true
		v.FgColor = gocui.ColorWhite
		fmt.Fprintln(v, "Line 1 in Right View")
		fmt.Fprintln(v, "Line 2 in Right View")
		fmt.Fprintln(v, "Line 3 in Right View")
	}

	// Initialisation une seule fois
	if !firstLaunch {
		if _, err := g.SetCurrentView("left"); err != nil {
			return err
		}

		firstLaunch = true
	}

	return nil
}

func keybindings(g *gocui.Gui) error {
	// Quit with Ctrl+C
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	// Cycle views with Tab
	if err := g.SetKeybinding("", 'l', gocui.ModNone, nextView); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'h', gocui.ModNone, prevView); err != nil {
		return err
	}

	// Move focus down with 'j'
	if err := g.SetKeybinding("", 'j', gocui.ModNone, cursorDown); err != nil {
		return err
	}

	// Move focus up with 'k'
	if err := g.SetKeybinding("", 'k', gocui.ModNone, cursorUp); err != nil {
		return err
	}

	return nil
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	log.Printf("Number of views: %d", len(views)) // This will also go to app.log
	if currentViewIndex >= len(views) {
		currentViewIndex = 0
		g.SetCurrentView(views[0])
	} else {
		currentViewIndex++
		g.SetCurrentView(views[currentViewIndex])
	}
	return nil
}

func prevView(g *gocui.Gui, v *gocui.View) error {
	if currentViewIndex <= 0 {
		currentViewIndex = len(views)
		g.SetCurrentView(views[currentViewIndex])
	} else {
		currentViewIndex++
		g.SetCurrentView(views[currentViewIndex])
	}
	return nil
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}
	cx, cy := v.Cursor()
	if err := v.SetCursor(cx, cy+1); err != nil {
		ox, oy := v.Origin()
		return v.SetOrigin(ox, oy+1)
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}
	cx, cy := v.Cursor()
	if err := v.SetCursor(cx, cy-1); err != nil {
		ox, oy := v.Origin()
		if oy > 0 {
			return v.SetOrigin(ox, oy-1)
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	Init()
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)
	g.SetCurrentView("left")

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
