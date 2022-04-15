package movieBrowser

type directories struct {
	Directories []*directory
}

func newDirectories() *directories {
	return new(directories)
}

func (d *directories) addDirectory(path string) error {
	if d.pathExists(path) {
		return pathDoesNotExistError{dir: path}
	}

	dir := newDirectory(path)
	dir.addFiles()
	d.Directories = append(d.Directories, dir)

	return nil
}

func (d *directories) pathExists(path string) bool {
	for k := range d.Directories {
		if d.Directories[k].Path == path {
			return true
		}
	}
	return false
}
