package toolbar

import (
	"erp/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ToolBar struct {
	Standard struct {
		Operation map[string]func() // 增删改查无效
		Print     map[string]func() // 打印导出EXCEL
		Paging    map[string]func()
		Quit      func()
	}
	Customize struct {
		ToolBar *widget.Toolbar
		icons   map[fyne.Resource]func()
	}
}

func (t *ToolBar) New(appType int, icons map[fyne.Resource]func()) fyne.CanvasObject {

	toolBar := t.mainToolBar(appType)
	subToolBar := widget.NewToolbar()

	for k, v := range icons {
		item := widget.NewToolbarAction(k, v)
		subToolBar.Append(item)
	}

	t.Customize.icons = icons
	t.Customize.ToolBar = subToolBar

	return container.NewVBox(
		toolBar,
		subToolBar,
	)
}

// mainToolBar 标准菜单
func (t *ToolBar) mainToolBar(appType int) fyne.CanvasObject {
	toolBar := widget.NewToolbar()
	switch appType {
	case data.Form:
		appendToolBarItem(toolBar, t.operation(nil, nil, nil, nil, nil)...)
		appendToolBarItem(toolBar, widget.NewToolbarSeparator())
		appendToolBarItem(toolBar, t.printToolBar(nil, nil)...)
		appendToolBarItem(toolBar, widget.NewToolbarSeparator())
		appendToolBarItem(toolBar, t.paging(nil, nil, nil, nil, nil)...)
		appendToolBarItem(toolBar, widget.NewToolbarSpacer())
		appendToolBarItem(toolBar, t.quit(nil)...)
	case data.Table:
		appendToolBarItem(toolBar, t.operation(nil, nil, nil, nil, nil)...)
		appendToolBarItem(toolBar, widget.NewToolbarSeparator())
		appendToolBarItem(toolBar, t.printToolBar(nil, nil)...)
		appendToolBarItem(toolBar, widget.NewToolbarSeparator())
		appendToolBarItem(toolBar, t.paging(nil, nil, nil, nil, nil)...)
		appendToolBarItem(toolBar, widget.NewToolbarSpacer())
		appendToolBarItem(toolBar, t.quit(nil)...)
	case data.Batch:
		appendToolBarItem(toolBar, t.quit(nil)...)
	case data.Para:
		appendToolBarItem(
			toolBar,
			widget.NewToolbarAction(theme.DocumentCreateIcon(), nil),
			widget.NewToolbarSeparator(),
			widget.NewToolbarAction(theme.InfoIcon(), nil),
			widget.NewToolbarAction(theme.LoginIcon(), func() {}),
		)
	case data.Query:
		appendToolBarItem(
			toolBar,
			widget.NewToolbarAction(theme.SearchIcon(), nil),
			widget.NewToolbarSeparator(),
		)
		appendToolBarItem(toolBar, widget.NewToolbarSpacer())
		appendToolBarItem(toolBar, t.quit(nil)...)
	case data.Report:
		appendToolBarItem(toolBar, t.quit(nil)...)
	default:
		appendToolBarItem(toolBar, t.quit(nil)...)
	}
	return toolBar
}

func (t *ToolBar) operation(i, r, m, q, u func()) []widget.ToolbarItem {
	t.Standard.Operation = map[string]func(){}
	t.Standard.Operation["insert"] = i
	t.Standard.Operation["remove"] = r
	t.Standard.Operation["modify"] = m
	t.Standard.Operation["query"] = q
	t.Standard.Operation["void"] = u

	return []widget.ToolbarItem{
		widget.NewToolbarAction(theme.ContentAddIcon(), i),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), r),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), m),
		widget.NewToolbarAction(theme.SearchIcon(), q),
		widget.NewToolbarAction(theme.DeleteIcon(), u),
	}
}
func (t *ToolBar) printToolBar(p, o func()) []widget.ToolbarItem {
	t.Standard.Print = map[string]func(){}
	t.Standard.Print["print"] = p
	t.Standard.Print["excel"] = o
	return []widget.ToolbarItem{
		widget.NewToolbarAction(theme.ListIcon(), o),
		widget.NewToolbarAction(theme.DocumentPrintIcon(), p),
	}
}
func (t *ToolBar) paging(f, p, j, n, l func()) []widget.ToolbarItem {
	t.Standard.Paging = map[string]func(){}
	t.Standard.Paging["first"] = f
	t.Standard.Paging["previous"] = p
	t.Standard.Paging["jump"] = j
	t.Standard.Paging["next"] = n
	t.Standard.Paging["last"] = l
	return []widget.ToolbarItem{
		widget.NewToolbarAction(theme.MediaFastRewindIcon(), f),
		widget.NewToolbarAction(theme.MediaSkipPreviousIcon(), p),
		widget.NewToolbarAction(theme.MediaRecordIcon(), j),
		widget.NewToolbarAction(theme.MediaSkipNextIcon(), n),
		widget.NewToolbarAction(theme.MediaFastForwardIcon(), l),
	}
}
func (t *ToolBar) quit(i func()) []widget.ToolbarItem {
	return []widget.ToolbarItem{
		widget.NewToolbarAction(theme.InfoIcon(), i),
		widget.NewToolbarAction(theme.LoginIcon(), func() {}),
	}
}

// appendToolBarItem 组成工具栏
func appendToolBarItem(t *widget.Toolbar, items ...widget.ToolbarItem) {
	for _, v := range items {
		t.Append(v)
	}
}
