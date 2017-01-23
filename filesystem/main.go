package filesystem


// Directory functions
/////////////////////////////////////////////////////////////////////////////
func LoadDir(path string) Directory {
	dir := Directory{}
	dir.Load(path)
	return dir
}
func ReadDir(path string) ([]Directory, []File) {
	dir := LoadDir(path)
	return dir.Read(false)
}
func ReadDirRecursive(path string) ([]Directory, []File) {
	dir := LoadDir(path)
	return dir.Read(true)
}
func CreateDir(path string) Directory {
	dir := Directory{}
	return dir.Create(path)
}
func CopyDir(source string, destination string, recurse bool, force bool) Directory {
	dir := LoadDir(source)
	return dir.Copy(destination, recurse, force)
}


// File functions
/////////////////////////////////////////////////////////////////////////////
func LoadFile(path string) File {
	file := File{}
	file.Load(path)
	return file
}

// Path functions
/////////////////////////////////////////////////////////////////////////////
func JoinPath(part ...string) string {
	return joinPath(part...)
}

func CheckPathExists(path string) bool {
	return checkPathExists(path)
}

func GetValidPath(path ...string) string {
	return getValidPath(path...)
}