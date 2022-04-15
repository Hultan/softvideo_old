package movieBrowser

import "fmt"

type pathDoesNotExistError struct {
	dir string
}

func (e pathDoesNotExistError) Error() string {
	return fmt.Sprintf("The dir '%s' does not exist", e.dir)
}
