package main

import "github.com/gotk3/gotk3/gtk"

func populateRecentMenu() {
	subMenu, _ := gtk.MenuNew()

	for k := range paths.paths {
		path := paths.paths[k]
		item, _ := gtk.MenuItemNewWithLabel(paths.paths[k])
		subMenu.Append(item)
		item.Connect("activate", func() {
			changePath(path)
		})
	}
	recentMenu.SetSubmenu(subMenu)
	recentMenu.ShowAll()
}
