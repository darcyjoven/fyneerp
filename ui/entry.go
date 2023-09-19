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
	"strconv"
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
	rec.SetMinSize(fyne.NewSize(200, 300))
	var data [][]string
	for i := 0; i < 100; i++ {
		s := strconv.Itoa(i)
		data = append(data, []string{s, s, s, s})
	}
	cir := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, object fyne.CanvasObject) {
			object.(*widget.Label).SetText(data[id.Row][id.Col])
		},
	)

	// Border 布局
	ccc := container.NewBorder(
		widget.NewLabel("www"),
		nil,
		nil,
		//, // 主要内容下
		nil,
		cir,
	)
	toolBar := container.NewVBox(
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
	)
	footer := container.NewHBox(
		container.NewVBox(
			layout.NewSpacer(),
			widget.NewLabel("staut is ok "),
		),
		container.NewVBox(
			layout.NewSpacer(),
			widget.NewLabel("staut is ok "),
		),
	)

	content := container.New(
		layout.NewBorderLayout(toolBar, footer, nil, nil),
		toolBar,
		footer,
		ccc,
	)
	return content
}

// SetMenu 设置菜单,可选项
func (app *Config) SetMenu(value map[string][]string, fn map[string]func()) {
	var m menu.Menu
	m.New(*app.win, value, fn)
	app.winCfg = &m
}
