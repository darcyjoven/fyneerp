package main

import (
	"erp/data"
	"erp/gui"
	theme2 "erp/theme"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&theme2.MyTheme{})
	w := a.NewWindow("test")

	cfg := gui.Config{}

	cfg.SetMenu(
		w,
		map[string][]string{
			"":   {"文件", "编辑", "视图", "工具", "关于"},
			"文件": {"打开", "保存", "另存为"},
			"编辑": {"复制", "粘贴", "剪切"},
			"视图": {"菜单栏", "状态栏"},
			"工具": {"设置", "切换视图"},
			"关于": {"注册信息", "注册", "版本信息"},
		},
		map[string]func(){},
	)
	cfg.SetToolBar(
		data.Form,
		widget.NewToolbarAction(theme.AccountIcon(), func() {}),
	)
	cfg.SetTop(container.NewGridWrap(fyne.NewSize(300, 40),
		container.NewBorder(
			nil, nil,
			widget.NewLabel("单号"),
			nil,
			widget.NewEntry(),
		),
		container.NewBorder(
			nil, nil,
			widget.NewLabel("单据日期"),
			nil,
			widget.NewEntry(),
		),
		container.NewBorder(
			nil, nil,
			widget.NewLabel("制单人"),
			nil,
			widget.NewEntry(),
		),
		container.NewBorder(
			nil, nil,
			widget.NewLabel("备注"),
			nil,
			widget.NewMultiLineEntry(),
		),
	))
	cfg.SetCenter(
		widget.NewTable(
			func() (int, int) {
				return 10, 10
			},
			func() fyne.CanvasObject {
				return widget.NewLabel("")
			},
			func(id widget.TableCellID, object fyne.CanvasObject) {
				object.(*widget.Label).SetText(strconv.Itoa(id.Col + id.Row))
			},
		),
	)
	str := binding.NewString()
	err := str.Set("test")
	if err != nil {
		return
	}
	cfg.SetFooter(str)
	c := cfg.MakeGUI()

	w.SetContent(c)

	w.Resize(fyne.NewSize(800, 800))
	w.ShowAndRun()
}
