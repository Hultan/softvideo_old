package movieBrowser

import "fmt"

type PathDoesNotExistError struct {
	dir string
}

func (e PathDoesNotExistError) Error() string {
	return fmt.Sprintf("The dir '%s' does not exist", e.dir)
}


