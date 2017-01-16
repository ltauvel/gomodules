package filesystem

import(
	"os"
	"io/ioutil"
	"github.com/ltauvel/gomodules/console"
)

type Directory struct {
	Item
}

func (directory *Directory) Load(path string) *Directory {
	console.PrintDebug("Loading directory", path)
	
	if directory.Exists(path) {
		d, _ := os.Stat(path)
		directory.Name = d.Name()
		directory.FullName = path
		directory.Size = d.Size()
		directory.Mode = d.Mode()
		directory.ModTime = d.ModTime()
		directory.Sys = d.Sys()
	} else {
		console.PrintError("File does ot exists")
	}
	
	return directory
}

func (directory Directory) Read() ([]Directory, []File) {
	var resultdirs []Directory
	var resultfiles []File

	console.PrintDebug("Reading", directory.FullName, "content")
	
	items, _ := ioutil.ReadDir(directory.FullName)
	for _, i := range items {

		ttfsitem := Item{
				Name: i.Name(),
				FullName: joinPath(directory.FullName, i.Name()),
				Size: i.Size(),
				Mode: i.Mode(),
				ModTime: i.ModTime(),
				Sys: i.Sys(),
			}
	
		if i.IsDir() {
			resultdirs = append(resultdirs, Directory{
				Item: ttfsitem,
			})
		} else {
			resultfiles = append(resultfiles, File{
				Item: ttfsitem,
			})
		}
	}
	
	return resultdirs, resultfiles
}

func (directory Directory) Create(path string) Directory {
	var result Directory
	
	fullpath := JoinPath(directory.FullName, path)
	
	console.PrintDebug("Creating directory", fullpath)

	// Create the directory
	os.MkdirAll(fullpath, 0777)
	
	// Assign result
	result.Load(fullpath)
	
	return result
}

func (directory Directory) Copy(destination string, recurse bool, force bool) Directory {
	var result Directory
	
	console.PrintDebug("Copying directory", directory.FullName, "to", destination)

	// Check if the file already exists
	if force || !directory.Exists(destination) {
	
		// Creating the backup directory
		os.MkdirAll(destination,0777)
		
		dirs, files := directory.Read()
		
		// Copy subdirectories
		if recurse {
			for _, dir := range dirs {
				dir.Copy(joinPath(destination, dir.Name), recurse, force)
			}
		}

		// Copy files
		for _, file := range files {
			file.Copy(joinPath(destination, file.Name), force)
		}

		// Assign result
		result.Load(destination)
		
	} else {
		console.PrintError("The specified destination file already exists.")
	}
	
	return result
}