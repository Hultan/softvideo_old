package movieBrowser

import (
	"os"
	"path"
	"path/filepath"
)

type directory struct {
	Name  string
	Path  string
	Files []*file
}

func newDirectory(directoryPath string) *directory {
	d := new(directory)
	d.Path = directoryPath
	d.Name = path.Base(directoryPath)
	return d
}

func (d *directory) addFiles() error {
	files, err := d.filePathWalkDir(d.Path)
	if err != nil {
		return err
	}

	for _, v := range files {
		file := newFile(v)
		d.Files = append(d.Files, file)
	}

	return nil
}

func (d *directory) filePathWalkDir(root string) ([]string, error) {
	var f []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			f = append(f, path)
		}
		return nil
	})
	return f, err
}
