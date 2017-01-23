package filesystem

import(
	"os"
	"io/ioutil"
	"github.com/ltauvel/gomodules/console"
)

type Directory struct {
	Item
}

func (directory *Directory) load(path string, requestbasepath string) {
	directory.Item.load(path, requestbasepath)
}

func (directory *Directory) Load(path string) {	
	directory.load(path, "")
}

func (directory Directory) read(recurse bool, requestbasepath string) ([]Directory, []File) {
	var resultdirs []Directory
	var resultfiles []File

	console.PrintDebug("Reading", directory.FullName, "content")
	
	items, _ := ioutil.ReadDir(directory.FullName)
	for _, i := range items {

		fsitem := Item{}
		fsitem.load(joinPath(directory.FullName, i.Name()), requestbasepath)
	
		if i.IsDir() {

			dir := Directory{
				Item: fsitem,
			}
		
			// Read subdirectories content
			if recurse {
				subdirs, subfiles := dir.read(recurse, requestbasepath)
				resultdirs = append(resultdirs, subdirs...)
				resultfiles = append(resultfiles, subfiles...)
			}
		
			resultdirs = append(resultdirs, dir)
		} else {
			resultfiles = append(resultfiles, File{
				Item: fsitem,
			})
		}
	}
	
	return resultdirs, resultfiles
}

func (directory Directory) Read(recurse bool) ([]Directory, []File) {
	if(recurse) {
		return directory.read(recurse, directory.FullName)
	} else {
		return directory.read(recurse, "")
	}
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
		
		dirs, files := directory.Read(false)
		
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