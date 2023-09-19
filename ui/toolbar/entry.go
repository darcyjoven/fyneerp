package toolbar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ToolBar struct {
	ToolBar *widget.Toolbar
	icons   map[fyne.Resource]func()
}

func (t *ToolBar) New(icons map[fyne.Resource]func()) fyne.CanvasObject {
	toolBar := widget.NewToolbar()

	for k, v := range icons {
		item := widget.NewToolbarAction(k, v)
		toolBar.Append(item)
	}

	t.icons = icons
	t.ToolBar = toolBar
	return toolBar
}
