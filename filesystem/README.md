# filesystem

`filesystem` is as a [Go](https://github.com/golang/go) package aimed at providing file system :

- Distinguished object types for directories and files
- Copy files or folders
- Path functions


Get directory information

	srcdir := filesystem.Directory{}
	srcdir.Load(srcpath)
	
	OR
	
	srcdir := filesystem.LoadDir(srcpath)

Copy a directory

	srcdir := filesystem.Directory{}
	srcdir.Load(srcpath)
	dstdir = srcdir.Copy(dstpath, true, true)
	
	OR
	
	srcdir := filesystem.LoadDir(srcpath)
	dstdir = srcdir.Copy(dstpath, true, true)
	
	OR
	
	dstdir = filesystem.CopyDir(srcpath, dstpath, true, true)


Read a directory

	srcdir := filesystem.Directory{}
	srcdir.Load(srcpath)
	subdirs, subfiles := srcdir.Read()
	
	OR
	
	srcdir := filesystem.LoadDir(srcpath)
	subdirs, subfiles := srcdir.Read()
	
	OR
	
	subdirs, subfiles := filesystem.ReadDir(srcpath)
	
	