package movieBrowser

import (
	"os"
	"path"
	"path/filepath"
)

type Directory struct {
	Name  string
	Path  string
	Files []*File
}

func NewDirectory(directoryPath string) *Directory {
	d := new(Directory)
	d.Path = directoryPath
	d.Name = path.Base(directoryPath)
	return d
}

func (d *Directory) AddFiles() error {
	files, err := d.filePathWalkDir(d.Path)
	if err != nil {
		return err
	}

	for _, v := range files {
		file := NewFile(v)
		d.Files = append(d.Files, file)
	}

	return nil
}

func (d *Directory) filePathWalkDir(root string) ([]string, error) {
	var f []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			f = append(f, path)
		}
		return nil
	})
	return f, err
}
