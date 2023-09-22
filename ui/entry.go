package ui

import (
	ddd "erp/data"
	"erp/ui/footer"
	"erp/ui/menu"
	"erp/ui/title"
	"erp/ui/toolbar"
	"erp/ui/top"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Config struct {
	win     *fyne.Window
	winCfg  *menu.Menu
	toolBar *toolbar.ToolBar
	footer  *footer.Footer
	top     *top.Top
}

// MakeUI UI设置
func (app *Config) MakeUI(win fyne.Window, appType int) fyne.CanvasObject {
	app.win = &win

	//	设置标题
	title.Title(win)

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
	ccc := container.New(
		layout.NewBorderLayout(
			app.top.New(),
			cir,
			nil,
			nil,
		),
		app.top.New(),
		nil, nil, nil,
		app.top.New(),
		cir,
		cir,
	)
	app.toolBar = &toolbar.ToolBar{
		Standard: struct {
			Operation map[string]func()
			Print     map[string]func()
			Paging    map[string]func()
			Quit      func()
		}{},
	}
	app.footer = &footer.Footer{}
	toolBar := app.toolBar.New(ddd.Form, nil)

	s := "123"
	ps := binding.BindString(&s)
	footer := app.footer.New(ps)

	content := container.New(
		layout.NewBorderLayout(toolBar, footer, nil, nil),
		toolBar,
		footer,
		ccc,
	)
	return content
}

// SetMenu 设置菜单,可选项-
func (app *Config) SetMenu(value map[string][]string, fn map[string]func()) {
	var m menu.Menu
	m.New(*app.win, value, fn)
	app.winCfg = &m
}
