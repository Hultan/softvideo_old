package main

type PathList struct {
	Paths []string
}

func NewPathList() *PathList {
	return new(PathList)
}

func (p *PathList) AddPath(path string) {
	if p.pathExists(path) {
		return
	}

	p.Paths = append(p.Paths, path)
}

func (p *PathList) pathExists(path string) bool {
	for k := range p.Paths {
		if p.Paths[k] == path {
			return true
		}
	}
	return false
}