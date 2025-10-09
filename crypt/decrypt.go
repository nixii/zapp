/*
 * Decrypt encrypted text with a key
 */
package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"fmt"
)

// Decrypt text
func Decrypt(ctxt []byte, mpwd string) ([]byte, error) {
	// default
	if len(ctxt) == 0 {
		return nil, nil
	}

	keyArray := sha256.Sum256([]byte(mpwd))
	key := keyArray[:]

	block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    if len(ctxt) < aes.BlockSize {
        return nil, fmt.Errorf("ciphertext too short")
    }

    iv := ctxt[:aes.BlockSize]
    ctxt = ctxt[aes.BlockSize:]

    mode := cipher.NewCBCDecrypter(block, iv)
    mode.CryptBlocks(ctxt, ctxt)

    // Unpad
    padding := int(ctxt[len(ctxt)-1])
    return ctxt[:len(ctxt)-padding], nil
}