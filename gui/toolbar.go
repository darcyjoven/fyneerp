package gui

import (
	"erp/data"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func myToolBar(items ...widget.ToolbarItem) *widget.Toolbar {
	if len(items) == 0 {
		return nil
	}
	return widget.NewToolbar(items...)
}

func getToolBar(appType int) *widget.Toolbar {
	switch appType {
	case data.Para:
		return toolBarPara()
	case data.Form:
		return toolBarForm()
	case data.Table:
		return toolBarTable()
	case data.Query:
		return toolBarQuery()
	case data.Batch:
		return toolBarBatch()
	case data.Report:
		return toolBarReport()
	default:
		return toolBarForm()
	}
}

func toolBarPara() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), nil),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.DocumentPrintIcon(), nil),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.InfoIcon(), nil),
		widget.NewToolbarAction(theme.LoginIcon(), nil),
	)
}
func toolBarForm() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), nil),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), nil),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), nil),
		widget.NewToolbarAction(theme.SearchIcon(), nil),
		widget.NewToolbarAction(theme.DeleteIcon(), nil),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ListIcon(), nil),
		widget.NewToolbarAction(theme.DocumentPrintIcon(), nil),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.MediaFastRewindIcon(), nil),
		widget.NewToolbarAction(theme.MediaSkipPreviousIcon(), nil),
		widget.NewToolbarAction(theme.MediaRecordIcon(), nil),
		widget.NewToolbarAction(theme.MediaSkipNextIcon(), nil),
		widget.NewToolbarAction(theme.MediaFastForwardIcon(), nil),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.InfoIcon(), nil),
		widget.NewToolbarAction(theme.LoginIcon(), nil),
	)
}

func toolBarTable() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentRemoveIcon(), nil),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), nil),
		widget.NewToolbarAction(theme.SearchIcon(), nil),
		widget.NewToolbarAction(theme.DeleteIcon(), nil),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ListIcon(), nil),
		widget.NewToolbarAction(theme.DocumentPrintIcon(), nil),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.MediaFastRewindIcon(), nil),
		widget.NewToolbarAction(theme.MediaSkipPreviousIcon(), nil),
		widget.NewToolbarAction(theme.MediaRecordIcon(), nil),
		widget.NewToolbarAction(theme.MediaSkipNextIcon(), nil),
		widget.NewToolbarAction(theme.MediaFastForwardIcon(), nil),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.InfoIcon(), nil),
		widget.NewToolbarAction(theme.LoginIcon(), nil),
	)
}

func toolBarQuery() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.SearchIcon(), nil),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ListIcon(), nil),
		widget.NewToolbarAction(theme.DocumentPrintIcon(), nil),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.MediaFastRewindIcon(), nil),
		widget.NewToolbarAction(theme.MediaSkipPreviousIcon(), nil),
		widget.NewToolbarAction(theme.MediaRecordIcon(), nil),
		widget.NewToolbarAction(theme.MediaSkipNextIcon(), nil),
		widget.NewToolbarAction(theme.MediaFastForwardIcon(), nil),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.InfoIcon(), nil),
		widget.NewToolbarAction(theme.LoginIcon(), nil),
	)
}
func toolBarBatch() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.SearchIcon(), nil),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.InfoIcon(), nil),
		widget.NewToolbarAction(theme.LoginIcon(), nil),
	)
}
func toolBarReport() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.SearchIcon(), nil),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.InfoIcon(), nil),
		widget.NewToolbarAction(theme.LoginIcon(), nil),
	)
}
