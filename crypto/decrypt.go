package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"io/ioutil"
	"os"
	"bytes"
)


func decrypt(filePath string, keystring string) bool {

	if isEncrypted(filePath) {
	
		// read content from your file
		ciphertext, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err.Error())
		}
		
		// Copy the encrypted bytes except the encryption verification prefix
		ciphertext = ciphertext[len(encryptpreffix):]
		
		// this is a key
		block, err := aes.NewCipher(formatKeystring(keystring))
		if err != nil {
			panic(err)
		}

		if len(ciphertext) < aes.BlockSize {
			panic("ciphertext too short")
		}
		
		iv := ciphertext[:aes.BlockSize]
		ciphertext = ciphertext[aes.BlockSize:]
		
		stream := cipher.NewCFBDecrypter(block, iv)
		stream.XORKeyStream(ciphertext, ciphertext)

		// Check that the content has successfully been decrypted
		if string(ciphertext[:len(decryptpreffix)]) == decryptpreffix {

			// Remove the existing file
			os.Remove(filePath)
			
			// create a new file for saving the encrypted data.
			writer, err := os.Create(filePath)
			if err != nil {
				panic(err.Error())
			}
			defer writer.Close()
			
			// Copy the decrypted bytes except the decryption verification prefix
			_, err = io.Copy(writer, bytes.NewReader(ciphertext[len(decryptpreffix):]))
			if err != nil {
				panic(err.Error())
			}
		} else {
			return false
		}
	}
	
	return true
}
