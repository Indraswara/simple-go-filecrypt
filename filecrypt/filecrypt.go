package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
	"os"
)

func padKey(key []byte) []byte {
	hash := sha256.Sum256(key)
	return hash[:]
}

func Encrypt(source string, password []byte) error {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		return err
	}

	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	plainText, err := io.ReadAll(srcFile)
	if err != nil {
		return err
	}

	key := padKey(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	cipherText := aesgcm.Seal(nil, nonce, plainText, nil)

	tempFile := source
	outFile, err := os.Create(tempFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if _, err = outFile.Write(nonce); err != nil {
		return err
	}
	if _, err = outFile.Write(cipherText); err != nil {
		return err
	}

	// Close the source file before deleting
	srcFile.Close()

	if err = os.Remove(source); err != nil {
		return err
	}

	if err = os.Rename(tempFile, source); err != nil {
		return err
	}

	return nil
}

func Decrypt(source string, password []byte) error {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		return err
	}

	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(srcFile, nonce); err != nil {
		return err
	}

	cipherText, err := io.ReadAll(srcFile)
	if err != nil {
		return err
	}

	key := padKey(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	plainText, err := aesgcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return err
	}

	tempFile := source
	outFile, err := os.Create(tempFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if _, err = outFile.Write(plainText); err != nil {
		return err
	}

	// Close the source file before deleting
	srcFile.Close()

	if err = os.Remove(source); err != nil {
		return err
	}

	if err = os.Rename(tempFile, source[:len(source)-4]); err != nil {
		return err
	}

	return nil
}
