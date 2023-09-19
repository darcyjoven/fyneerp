package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"erp/ui"
)

func main() {
	a := app.New()
	w := a.NewWindow("test")

	var cfg ui.Config
	w.SetContent(cfg.MakeUI(w, ui.Form))
	w.Resize(fyne.NewSize(800, 800))
	w.ShowAndRun()
}
