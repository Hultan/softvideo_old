package main

import "github.com/gotk3/gotk3/gtk"

func populateRecentMenu() {
	subMenu, _ := gtk.MenuNew()

	for k := range pathList.Paths {
		item, _ := gtk.MenuItemNewWithLabel(pathList.Paths[k])
		subMenu.Append(item)
		item.Connect("activate", func() {
			changePath(pathList.Paths[k])
		})
	}
	recentMenu.SetSubmenu(subMenu)
	recentMenu.ShowAll()
}
