package movieBrowser

type Directories struct {
	Directories []*Directory
}

func NewDirectories() *Directories {
	return new(Directories)
}

func (d *Directories) AddDirectory(path string) error {
	if d.pathExists(path) {
		return PathDoesNotExistError{dir: path}
	}

	dir := NewDirectory(path)
	dir.AddFiles()
	d.Directories = append(d.Directories, dir)

	return nil
}

func (d *Directories) pathExists(path string) bool {
	for k := range d.Directories {
		if d.Directories[k].Path == path {
			return true
		}
	}
	return false
}
