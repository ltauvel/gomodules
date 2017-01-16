package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"
	"bytes"
)

func isEncrypted(filePath string) bool {
	// read content from your file
    content, err := ioutil.ReadFile(filePath)
    if err != nil {
        panic(err.Error())
    }
	
	if string(content[:len(encryptpreffix)]) == encryptpreffix {
		return true
	}
	
	return false
}

func encrypt(filePath string, keystring string) bool {

	if !isEncrypted(filePath) {

		// read content from your file
		plaintext, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err.Error())
		}
		
		// Copy the plain text bytes adding the decryption verification prefix
		plaintext = prePend(decryptpreffix, plaintext)
		
		// this is a key
		block, err := aes.NewCipher(formatKeystring(keystring))
		if err != nil {
			panic(err)
		}

		// The IV needs to be unique, but not secure. Therefore it's common to
		// include it at the beginning of the ciphertext.
		ciphertext := make([]byte, aes.BlockSize+len(plaintext))
		iv := ciphertext[:aes.BlockSize]
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			panic(err)
		}

		stream := cipher.NewCFBEncrypter(block, iv)
		stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

		// Remove the existing file
		os.Remove(filePath)
		
		// create a new file for saving the encrypted data.
		writer, err := os.Create(filePath)
		if err != nil {
			panic(err.Error())
		}
		defer writer.Close()
		
		// Copy the encrypted bytes with the encryption verification prefix
		_, err = io.Copy(writer, bytes.NewReader(prePend(encryptpreffix, ciphertext)))
		if err != nil {
			panic(err.Error())
		}
		
		return true
    }
	
	return false
}
