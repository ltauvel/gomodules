package targz

func Extract(archivepath string, targetpath string) {
	extract(archivepath, targetpath)
}

func CompressFolder(targetpath string, folderpath string) {
	compressFolder(targetpath, folderpath)
}