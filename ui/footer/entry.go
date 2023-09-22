package footer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Footer struct {
	bindStrings []binding.String
}

// New 新建页脚
func (f *Footer) New(bindString ...binding.String) fyne.CanvasObject {
	f.bindStrings = bindString

	content := container.New(layout.NewHBoxLayout())
	for _, v := range bindString {
		content.Add(widget.NewLabelWithData(v))
	}
	return content
}
