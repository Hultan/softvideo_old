package movieBrowser

import (
	"path"
)

type File struct {
	Name     string
	Path     string
	Favorite bool
}

func NewFile(filePath string) *File {
	f := new(File)
	f.Path = filePath
	f.Name = path.Base(filePath)
	return f
}
