package top

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Top struct {
}

func (t *Top) New() fyne.CanvasObject {
	content := container.NewGridWithColumns(10)
	for i := 0; i <= 20; i++ {
		content.Add(widget.NewLabel(strconv.Itoa(i)))
		content.Add(widget.NewEntry())
	}
	return content
}
