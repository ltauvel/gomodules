package filesystem

import(
	"os"
	"io"
	"github.com/ltauvel/gomodules/console"
)

type File struct {
	Item
}

func (file *File) Load(path string) *File {
	console.PrintDebug("Loading file", path)
	
	if file.Exists(path) {
		f, _ := os.Stat(path)
		file.Name = f.Name()
		file.FullName = path
		file.Size = f.Size()
		file.Mode = f.Mode()
		file.ModTime = f.ModTime()
		file.Sys = f.Sys()
	} else {
		console.PrintError("File does ot exists")
	}
	
	return file
}

func (file File) Copy(destination string, force bool) File {
	var result File

	// Check if the file already exists
	if force || !file.Exists(destination) {
	
		console.PrintDebug("Copying file", file.FullName, "to", destination)
	
		// Open the source file
		reader, err := os.Open(file.FullName)
		if err != nil {
			panic(err)
		}
		defer reader.Close()
		
		// Write the destination file
		writer, err := os.Create(destination)
		if err != nil {
			panic(err)
		}
		if _, err := io.Copy(writer, reader); err != nil {
			writer.Close()
			panic(err)
		}
		defer writer.Close()
		
		// Assign result
		result.Load(destination)
		
	} else {
		console.PrintError("The specified destination file already exists.")
	}
	
	return result
}
