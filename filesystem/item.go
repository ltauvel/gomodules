package filesystem

import(
	"os"
	"time"
)

type Item struct {
	Name string
	FullName string	
	Size int64
	Mode os.FileMode
	ModTime time.Time
	Sys interface{}
}

func (item *Item) Exists(path string) bool {
    return checkPathExists(path)
}