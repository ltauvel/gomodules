package filesystem

import(
	"os"
	"time"
	"strings"
	"regexp"
	"github.com/ltauvel/gomodules/console"
)

type Item struct {
	baserequestpath string
	BaseFullName string
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

func (item *Item) RelativeName() string {
	var r = regexp.MustCompile(`^(\\|/)`)

    return r.ReplaceAllString(strings.Replace(item.FullName, item.baserequestpath, "", -1), "")
}

func (item *Item) load(path string, requestbasepath string) {
	if item.Exists(path) {
		console.PrintDebug("Loading item", path)
		
		var r = regexp.MustCompile(`(\\|/)$`)
		
		i, _ := os.Stat(path)
		item.Name =	i.Name()
		item.FullName = path
		item.Size = i.Size()
		item.Mode = i.Mode()
		item.ModTime = i.ModTime()
		item.Sys = i.Sys()
		item.BaseFullName = r.ReplaceAllString(strings.Replace(item.FullName, item.Name, "", -1), "")
		if len(requestbasepath) > 0 {
			item.baserequestpath = r.ReplaceAllString(requestbasepath, "")
		} else {
			item.baserequestpath = item.BaseFullName
		}
	} else {
		console.PrintError("item does not exists")
	}
}