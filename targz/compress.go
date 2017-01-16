package targz

import (
	"os"
	"archive/tar"
	"log"
	"io"
	"strings"
	"compress/gzip"
	"path/filepath"
	"github.com/ltauvel/gomodules/console"
	"github.com/ltauvel/gomodules/filesystem"
)

func addDir(tarfileWriter * tar.Writer, path string, relativepath string, recurse bool) {
	console.PrintDebug("Adding directory ./", relativepath)

	// Opening directory to compress
	dir, err := os.Open(path)
	if err != nil {			
		log.Fatalln(err)
	}
	defer dir.Close()

	// Grabbing all the files of the directory to compress
	files, err := dir.Readdir(0)
	if err != nil {			
		log.Fatalln(err)
	}
	
	// Looping on each item in the directory to compress
	for _, fileInfo := range files {

		if fileInfo.IsDir() && recurse {
			addDir(tarfileWriter, filesystem.JoinPath(path, fileInfo.Name()), relativepath + fileInfo.Name() + "/", true)
		} else {
			console.PrintDebug("Adding file", fileInfo.Name())
			// see https://www.socketloop.com/tutorials/go-file-path-independent-of-operating-system

			file, err := os.Open(dir.Name() + string(filepath.Separator) + fileInfo.Name())
			if err != nil {
				log.Fatalln(err)
			}
			defer file.Close()

			// Prepare the TAR header
			header := new(tar.Header)
			header.Name = relativepath + filepath.Base(file.Name())
			header.Size = fileInfo.Size()
			header.Mode = int64(fileInfo.Mode())
			header.ModTime = fileInfo.ModTime()
			header.Typeflag = tar.TypeReg


			if err = tarfileWriter.WriteHeader(header) ; err != nil {
				log.Fatalln(err)
			}

			if _, err = io.Copy(tarfileWriter, file) ; err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func compressFolder(targetpath string, folderpath string) {
	console.PrintDebug("Creating archive", targetpath, "for folder", folderpath)

		// Create a destination file
		tarfile, err := os.Create(targetpath)
		if err != nil {			
			log.Fatalln(err)
		}
		defer tarfile.Close()
		
		// Create the GZIP writer if user add .gz in the destination filename
		var fileWriter io.WriteCloser = tarfile
		if strings.HasSuffix(targetpath, ".gz") {
			fileWriter = gzip.NewWriter(tarfile)
			defer fileWriter.Close()
		}

		// Create the TAR writer
		tarfileWriter := tar.NewWriter(fileWriter)
		defer tarfileWriter.Close()

		// Looping on each item in the directory to compress
		addDir(tarfileWriter, folderpath, "", true)
}
