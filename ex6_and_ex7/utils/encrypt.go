package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(encoded_string []byte) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, []byte(MySecret)[:block.BlockSize()])
	cipherText := make([]byte, len(encoded_string))
	cfb.XORKeyStream(cipherText, encoded_string)
	return Encode(cipherText), nil
}

func Decode(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText, err := Decode(text)
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBDecrypter(block, []byte(MySecret)[:block.BlockSize()])
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
