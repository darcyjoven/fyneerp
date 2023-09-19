package ui

import (
	"erp/ui/menu"
	"erp/ui/title"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type Config struct {
	win    *fyne.Window
	winCfg *menu.Menu
}

// MakeUI UI设置
func (app *Config) MakeUI(win fyne.Window, appType int) fyne.CanvasObject {
	app.win = &win

	//	设置标题
	title.Title(win)

	// ToolBar 设置标准ToolBar
	rec := canvas.NewRectangle(color.Black)
	cir := canvas.NewRectangle(color.Opaque)
	rec.SetMinSize(fyne.NewSize(200, 300))
	cir.SetMinSize(fyne.NewSize(200, 300))

	content := container.NewStack(
		container.New(
			layout.NewVBoxLayout(),
			container.NewVBox(
				container.NewHBox(
					widget.NewButton("ok", func() {}),
					widget.NewButton("ok", func() {}),
					widget.NewButton("ok", func() {}),
					widget.NewButton("ok", func() {}),
				), // toolBar
				container.NewGridWrap(
					fyne.NewSize(100, 20),
					widget.NewButton("ok", func() {}),
					widget.NewButton("ok", func() {}),
					widget.NewButton("ok", func() {}),
					widget.NewButton("ok", func() {}),
					widget.NewButton("ok", func() {}),
				), // DIT toolBar
			),
			container.NewVSplit(
				rec, // 主要内容上
				cir, // 主要内容下
			),
			container.NewHBox(
				widget.NewLabel("staut is ok "),
				widget.NewLabel("staut is not ok "),
			), // footer
		),
	)

	return content
}

// SetMenu 设置菜单,可选项
func (app *Config) SetMenu(value map[string][]string, fn map[string]func()) {
	var m menu.Menu
	m.New(*app.win, value, fn)
	app.winCfg = &m
}
