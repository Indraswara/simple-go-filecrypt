package main

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	funciton := os.Args[1]

	switch funciton {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandler()
	case "decrypt":
		decryptHandler()

	default:
		fmt.Println("Run encrypt to enxrypt a file or decrypt to decrypt a file. Run help to see all the commands.")
		os.Exit(0)
	}
}

func printHelp() {
	fmt.Println("Usage: ")
	fmt.Println("\tgo run . encrypt /path-to-file")
	fmt.Println("\tgo run . decrypt /path-to-file")

	fmt.Println("Commands:")
	fmt.Println("\tencrypt - encrypts a file")
	fmt.Println("\tdecrypt - decrypts a file")
	fmt.Println("help - prints this help message")
}

func encryptHandler() {
	if len(os.Args) < 3 {
		fmt.Println("Missing file path. Run help to see all the commands.")
		os.Exit(0)
	}
	filepath := os.Args[2]
	if !validateFile(filepath) {
		panic("File does not exist")
	}
	password := getPassword()
	fmt.Println("Encrypting file")
	filecrypt.Encrypt(file, password)
	fmt.Println("File successfully encrypted")
}

func decryptHandler() {
	if len(os.Args) < 3 {
		fmt.Println("Missing file path. Run help to see all the commands.")
		os.Exit(0)
	}
	filepath := os.Args[2]
	if !validateFile(filepath) {
		panic("File does not exist")
	}

	fmt.Println("Enter Password: ")
	password, _ := term.ReadPassword(0)
	fmt.Println("Decrypting file")
	filecrypt.Decrypt(file, password)
	fmt.Println("File successfully Decrypted")
}

func getPassword() []byte {

	fmt.Print("Enter Password: ")
	password, _ := term.ReadPassword(0)
	fmt.Print("\nConfirm Password: ")
	confirmPassword, _ := term.ReadPassword(0)
	if !validatePassword(password, confirmPassword) {
		fmt.Println("Passwords do not match")
		return getPassword()
	}
	return password
}

func validatePassword(pass1 []byte, pass2 []byte) bool {
	if !bytes.Equal(pass1, pass2) {
		return false
	}
	return true
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
