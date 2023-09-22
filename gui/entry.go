package gui

import (
	"erp/data"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

// Config GUI基础配置
type Config struct {
	UI struct {
		ToolBar *fyne.Container
		Footer  *fyne.Container
		Top     *fyne.Container
		Center  *widget.Table
		Bottom  *widget.Table
		Menu    *fyne.MainMenu
	}
	Data struct {
		ToolBar map[string]func()
	}
}

// MakeGUI 窗口初始化
func (app *Config) MakeGUI() fyne.CanvasObject {
	// 检查窗口子项目初始化是否完毕

	// ToolBar
	// Footer
	// Content
	// Menu

	return container.New(
		layout.NewBorderLayout(
			app.UI.ToolBar,
			app.UI.Footer,
			nil, nil),
		app.UI.ToolBar,
		app.UI.Footer,
		getContent(app.UI.Top, app.UI.Center, app.UI.Bottom),
	)
}

// SetToolBar 设置工具栏
func (app *Config) SetToolBar(appType int, toolBarItems ...widget.ToolbarItem) {
	app.UI.ToolBar = container.NewVBox(
		getToolBar(appType),
		myToolBar(toolBarItems...),
	)
}

// SetFooter 设置底部栏,只需要设置绑定的数据即可
func (app *Config) SetFooter(binds ...binding.String) {
	app.UI.Footer = container.NewHBox()
	for _, v := range binds {
		app.UI.Footer.Add(widget.NewLabelWithData(v))
	}
}

// SetMenu 设置菜单
func (app *Config) SetMenu(win fyne.Window, value map[string][]string, fn map[string]func()) {
	mainMenu := fyne.NewMainMenu()
	for _, v := range value[""] {
		menu := fyne.NewMenu(v)
		menu.Items = expandMenu(v, value, fn)
		mainMenu.Items = append(mainMenu.Items, menu)
	}
	app.UI.Menu = mainMenu
	win.SetMainMenu(mainMenu)
}

// SetMenuAction 设置菜单Action
func (app *Config) SetMenuAction(id string, fn func()) {
	// 	expand menu
	for _, v := range app.UI.Menu.Items {
		menu, ok := getMenu(v, id)
		if ok {
			menu.Action = fn
			break
		}
	}
}

// SetTop 设置TOP
func (app *Config) SetTop(object *fyne.Container) {
	app.UI.Top = object
}

// SetCenter 设置主要内容
func (app *Config) SetCenter(table *widget.Table) {
	app.UI.Center = table
}

// SetBottom 设置底部内容
func (app *Config) SetBottom(table *widget.Table) {
	app.UI.Bottom = table
}

// test 测试UI
func (app *Config) test() {
	app.UI.ToolBar = container.NewVBox(
		getToolBar(data.Form),
		myToolBar(
			widget.NewToolbarAction(theme.DocumentCreateIcon(), nil),
			widget.NewToolbarSeparator(),
			widget.NewToolbarAction(theme.DocumentPrintIcon(), nil),
			widget.NewToolbarSpacer(),
			widget.NewToolbarAction(theme.InfoIcon(), nil),
			widget.NewToolbarAction(theme.LoginIcon(), nil),
		),
	)
	app.UI.Footer = container.NewHBox(
		widget.NewLabel("Error"),
		widget.NewLabel("Success"),
	)
	app.UI.Top = container.NewGridWithColumns(10)
	for i := 0; i < 20; i++ {
		app.UI.Top.Add(widget.NewLabel(strconv.Itoa(i)))
		app.UI.Top.Add(widget.NewEntry())
	}
	var data1 [][]int
	for i := 0; i <= 50; i++ {
		data1 = append(data1, []int{1, 2, 3, 4})
	}
	app.UI.Center = widget.NewTable(
		func() (int, int) {
			return len(data1), len(data1[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, object fyne.CanvasObject) {
			object.(*widget.Label).SetText(strconv.Itoa(data1[id.Row][id.Col]))
		},
	)
	app.UI.Bottom = widget.NewTable(
		func() (int, int) {
			return len(data1), len(data1[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, object fyne.CanvasObject) {
			object.(*widget.Label).SetText(strconv.Itoa(data1[id.Row][id.Col] * 100))
		},
	)
}

// Demo 是

// Deprecated: Use container.NewStack() instead.
func (app *Config) Demo() fyne.CanvasObject {
	// 三段
	a := container.NewStack(
		widget.NewLabel("test"),
		widget.NewLabel("test"),
	)
	var data1 [][]int
	for i := 0; i <= 50; i++ {
		data1 = append(data1, []int{1, 2, 3, 4})
	}
	b := widget.NewTable(
		func() (int, int) {
			return len(data1), len(data1[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, object fyne.CanvasObject) {
			object.(*widget.Label).SetText(strconv.Itoa(data1[id.Row][id.Col]))
		},
	)
	c := widget.NewTable(
		func() (int, int) {
			return len(data1), len(data1[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.TableCellID, object fyne.CanvasObject) {
			object.(*widget.Label).SetText(strconv.Itoa(data1[id.Row][id.Col] * 10))
		},
	)
	d := container.NewVSplit(b, c)
	return container.NewBorder(a, nil, nil, nil, d)
}
