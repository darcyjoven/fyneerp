package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// expand 展开菜单
func expandMenu(menuValue string, value map[string][]string, fn map[string]func()) []*fyne.MenuItem {
	var menuItems []*fyne.MenuItem
	for _, v := range value[menuValue] {
		item := fyne.NewMenuItem(v, func() {})
		_, ok := value[v]
		if ok {
			// 如果还有下级菜单
			menu := fyne.NewMenu(v, expandMenu(v, value, fn)...)
			item.ChildMenu = menu
		} else {
			// 如果没有下级菜单
			item.Action = fn[v]
		}
		menuItems = append(menuItems, item)
	}
	return menuItems
}

// getContent 组成主要内容容器
func getContent(t *fyne.Container, c, b *widget.Table) *fyne.Container {

	content := new(fyne.Container)

	if c == nil && b == nil {
		// 只有一个内容
		content = container.NewStack(t)
	} else if b == nil {
		// 	只有两个内容
		content = container.NewBorder(
			t, nil, nil, nil,
			c,
		)
	} else {
		// 	三个内容都存在
		temp := container.NewVSplit(c, b)
		content = container.NewBorder(
			t, nil, nil, nil,
			temp,
		)
	}
	return content
}

// getMenu 用递归的方式查询指定名称的菜单
func getMenu(menu *fyne.Menu, id string) (*fyne.MenuItem, bool) {
	for _, v := range menu.Items {
		if v.ChildMenu != nil {
			// 	有下阶菜单的时候，进行递归查询
			return getMenu(v.ChildMenu, id)
		} else {
			if v.Label == id {
				return v, true
			}
		}
	}
	return nil, false
}
