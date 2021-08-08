package main

import "github.com/gotk3/gotk3/gtk"

func populateRecentMenu() {
	subMenu, _ := gtk.MenuNew()

	for k := range pathList.Paths {
		path := pathList.Paths[k]
		item, _ := gtk.MenuItemNewWithLabel(pathList.Paths[k])
		subMenu.Append(item)
		item.Connect("activate", func() {
			changePath(path)
		})
	}
	recentMenu.SetSubmenu(subMenu)
	recentMenu.ShowAll()
}
