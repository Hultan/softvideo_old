package movieBrowser

import (
	"path"
)

type file struct {
	Name     string
	Path     string
	Favorite bool
}

func newFile(filePath string) *file {
	f := new(file)
	f.Path = filePath
	f.Name = path.Base(filePath)
	return f
}
