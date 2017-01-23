package crypto

func IsEncrypted(filePath string) bool {
	return isEncrypted(filePath)
}

func Encrypt(filePath string, keystring string) bool {
	return encrypt(filePath, keystring) 
}

func Decrypt(filePath string, keystring string) bool {
	return decrypt(filePath, keystring) 
}

func GetChecksum(filePath string) []byte {
	return getChecksum(filePath, false)
}

func GetDirectoryChecksum(filePath string, ignoreemptydir bool, ignoreitems ...string) []byte {
	return getChecksum(filePath, ignoreemptydir, ignoreitems...)
}