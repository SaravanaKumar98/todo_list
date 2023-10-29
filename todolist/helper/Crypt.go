package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"
)

//crypt

func encrypt(text string, key []byte) (string, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintext := []byte(text)

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(text string, key []byte) (string, error) {

	ciphertext, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

func Crypt(originalText string, need string) (string, error) {

	key := []byte(os.Getenv("KEY"))

	if need == "encrypt" {
		str, err := encrypt(originalText, key)
		if err != nil {
			return "", err
		}

		return str, nil
	}

	str, err := decrypt(originalText, key)
	if err != nil {
		return "", err
	}
	return str, nil

}
