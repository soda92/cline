package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"github.com/UserExistsError/conpty"
)

func eval(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("germ")

	p, err := conpty.Start(`pwsh.exe`)
	eval(err)

	defer p.Close()

	// os.Setenv("TERM", "xterm-256color")
	terminal := NewTerminal(p)

	w.SetContent(
		container.New(
			layout.NewGridWrapLayout(fyne.NewSize(630, 630)),
			terminal,
		),
	)
	w.Canvas().Focus(terminal)

	w.ShowAndRun()
}
