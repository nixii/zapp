/*
 * Handle encrypting a string into bytes with a master password
 */
package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

// Encrypt bytes to other bytes
func Encrypt(txt []byte, mpwd string) ([]byte, error) {
	
	// Get the key
	keyArray := sha256.Sum256([]byte(mpwd))
	key := keyArray[:]

	// Load the cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// The padding required to make the text correct
	padding := aes.BlockSize - len(txt) % aes.BlockSize
	padtext := append(txt, bytes.Repeat([]byte{byte(padding)}, padding)...)

	// Prepare the cipher text
	ciphertext := make([]byte, aes.BlockSize + len(padtext))
	iv := ciphertext[:aes.BlockSize]

	// Read the iv
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Encrypt the data
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], padtext)

	// Return the encrypted data
	return ciphertext, nil
}