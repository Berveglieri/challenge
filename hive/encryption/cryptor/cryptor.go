package cryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)

func CreateHash(key string) string {
	hasher := sha256.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))[:32]
}

func Encrypt(filename string, passphrase string, destFile string) []byte {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher([]byte(CreateHash(passphrase)))
	if err != nil{
		panic(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(data), nil)

	ioutil.WriteFile(destFile, ciphertext, 0755)
	return ciphertext
}

func Decrypt(filename string, passphrase string, destFile string) []byte {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	key := []byte(CreateHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		panic(err.Error())
	}

	ioutil.WriteFile(destFile, plaintext, 0755)
	return plaintext
}

