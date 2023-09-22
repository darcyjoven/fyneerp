package toolbar

import (
	"erp/data"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
)

func TestToolBar_New(t *testing.T) {
	a := test.NewApp()
	w := a.NewWindow("tess")
	toolBar := new(ToolBar)

	w.SetContent(toolBar.New(
		data.Form,
		map[fyne.Resource]func(){
			theme.ContentAddIcon():     func() {},
			theme.DocumentCreateIcon(): func() {},
			theme.SearchIcon():         func() {},
			theme.DeleteIcon():         func() {},
		},
	))
}
