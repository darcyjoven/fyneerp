package menu

import "fyne.io/fyne/v2"

type Menu struct {
	MainMenu *fyne.MainMenu
	fn       map[string]func()
	value    map[string][]string
}

// New 创建菜单
func (m *Menu) New(win fyne.Window, value map[string][]string, fn map[string]func()) {
	mainMenu := fyne.NewMainMenu()
	for _, v := range value[""] {
		menu := fyne.NewMenu(v)
		menu.Items = expand(v, value, fn)
		mainMenu.Items = append(mainMenu.Items, menu)
	}
	m.fn = fn
	m.value = value
	m.MainMenu = mainMenu
	win.SetMainMenu(mainMenu)
}

// expand 展开菜单
func expand(menuValue string, value map[string][]string, fn map[string]func()) []*fyne.MenuItem {
	var menuItems []*fyne.MenuItem
	for _, v := range value[menuValue] {
		item := fyne.NewMenuItem(v, func() {})
		_, ok := value[v]
		if ok {
			// 如果还有下级菜单
			menu := fyne.NewMenu(v, expand(v, value, fn)...)
			item.ChildMenu = menu
		} else {
			// 如果没有下级菜单
			item.Action = fn[v]
		}
		menuItems = append(menuItems, item)
	}
	return menuItems
}
