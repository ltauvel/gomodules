package targz

import (
	"os"
	"archive/tar"
	"io"
	"strings"
	"compress/gzip"
	"fmt"
	"github.com/ltauvel/gomodules/console"
	"github.com/ltauvel/gomodules/filesystem"
)

func extract(archivepath string, targetpath string) {

	file, err := os.Open(archivepath)
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	defer file.Close()

	var fileReader io.ReadCloser = file

	// just in case we are reading a tar.gz file, add a filter to handle gzipped file
	if strings.HasSuffix(archivepath, ".gz") {
		if fileReader, err = gzip.NewReader(file); err != nil {
			panic(err)
			os.Exit(1)
		}
		defer fileReader.Close()
	}

	tarBallReader := tar.NewReader(fileReader)

	// Extracting tarred files
	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
			os.Exit(1)
		}

		// get the individual filename and extract to the current directory
		filename := filesystem.JoinPath(targetpath, header.Name)


		switch header.Typeflag {
			case tar.TypeDir:
				// handle directory
				console.PrintDebug("Creating directory", filename)
				err = os.MkdirAll(filename, os.FileMode(header.Mode)) // or use 0755 if you prefer

				if err != nil {
					panic(err)
					os.Exit(1)
				}

			case tar.TypeReg:
				// handle normal file
				console.PrintDebug("Extracting", filename)
				
				// Create the subdirectoy if needed
				os.MkdirAll(filename, os.FileMode(header.Mode)) 
				os.RemoveAll(filename) 
				
				
				writer, err := os.Create(filename)
				if err != nil {
					panic(err)
					os.Exit(1)
				}
				defer writer.Close()

				io.Copy(writer, tarBallReader)

				err = os.Chmod(filename, os.FileMode(header.Mode))
				if err != nil {
					panic(err)
					os.Exit(1)
				}
				
			default:
				fmt.Printf("Unable to untar type : %c in file %s", header.Typeflag, filename)
		}
	 }
}