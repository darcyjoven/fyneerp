package toolbar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"testing"
)

func TestToolBar_New(t *testing.T) {
	a := test.NewApp()
	w := a.NewWindow("tess")
	toolBar := new(ToolBar)

	w.SetContent(toolBar.New(
		map[fyne.Resource]func(){
			theme.ContentAddIcon():     func() {},
			theme.DocumentCreateIcon(): func() {},
			theme.SearchIcon():         func() {},
			theme.DeleteIcon():         func() {},
		},
	))
}
