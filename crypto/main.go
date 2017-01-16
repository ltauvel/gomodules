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