package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// 这是一个三段式布局，上下为固定宽度，中间占用剩余空间
type threeLayout struct {
	horizontal   bool
	topHeight    float64
	bottomHeight float64
	top          fyne.CanvasObject
	center       fyne.CanvasObject
	bottom       fyne.CanvasObject
}

func NewHThreeLayout(top, center, bottom fyne.CanvasObject) fyne.Layout {
	return &threeLayout{
		horizontal: true,
		top:        top,
		center:     center,
		bottom:     bottom,
	}
}

func NewVBoxLayout(top, center, bottom fyne.CanvasObject) fyne.Layout {
	return &threeLayout{
		horizontal: false,
		top:        top,
		center:     center,
		bottom:     bottom,
	}
}
func (g *threeLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	miniSize := fyne.NewSize(0, 0)

	return miniSize
}
func (g *threeLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	padding := theme.Padding()
	var topSize, bottomSize, centerSize fyne.Size
	if g.top != nil && g.top.Visible() {
		topHeight := g.top.MinSize().Height
		if topHeight == 0 {
			topHeight = 100
		}
		g.top.Resize(fyne.NewSize(size.Width, topHeight))
		g.top.Move(fyne.NewPos(0, 0))
		topSize = fyne.NewSize(size.Width, topHeight+padding)
	}

}

func (g *threeLayout) SetSize(top, bottom float64) {
	g.topHeight = top
	g.bottomHeight = bottom
}
