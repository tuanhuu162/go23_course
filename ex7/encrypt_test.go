package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/tuanhuu162/go23_course/ex6/utils"
	"testing"
)

func TestEncrypt(t *testing.T) {
	encrypted, err := utils.Encrypt([]byte("test"))
	assert.Nil(t, err)
	assert.NotEqual(t, encrypted, "test")
}

func TestDecrypt(t *testing.T) {
	encrypted, err := utils.Encrypt([]byte("test"))
	assert.Nil(t, err)
	decrypted, err := utils.Decrypt(encrypted)
	assert.Nil(t, err)
	assert.Equal(t, decrypted, "test")
}
