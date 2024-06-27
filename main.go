package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/Indraswara/file-encrypt/filecrypt"
	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandler()
	case "decrypt":
		decryptHandler()
	default:
		fmt.Println("Run 'encrypt' to encrypt a file or 'decrypt' to decrypt a file. Run 'help' to see all the commands.")
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
	fmt.Println("\thelp - prints this help message")
}

func encryptHandler() {
	if len(os.Args) < 3 {
		fmt.Println("Missing file path. Run help to see all the commands.")
		os.Exit(0)
	}
	filepath := os.Args[2]
	if !validateFile(filepath) {
		fmt.Println("File does not exist")
		os.Exit(1)
	}
	password := getPassword()
	fmt.Println("\nEncrypting file")
	filecrypt.Encrypt(filepath, password)
	fmt.Println("File successfully encrypted")
}

func decryptHandler() {
	if len(os.Args) < 3 {
		fmt.Println("Missing file path. Run help to see all the commands.")
		os.Exit(0)
	}
	filepath := os.Args[2]
	if !validateFile(filepath) {
		fmt.Println("File does not exist")
		os.Exit(1)
	}

	fmt.Print("Enter Password: ")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error reading password:", err)
		os.Exit(1)
	}
	fmt.Println("\nDecrypting file")
	filecrypt.Decrypt(filepath, password)
	fmt.Println("File successfully decrypted")
}

func getPassword() []byte {
	fmt.Print("Enter Password: ")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error reading password:", err)
		os.Exit(1)
	}
	fmt.Print("\nConfirm Password: ")
	confirmPassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error reading password:", err)
		os.Exit(1)
	}
	if !validatePassword(password, confirmPassword) {
		fmt.Println("Passwords do not match")
		return getPassword()
	}
	return password
}

func validatePassword(pass1 []byte, pass2 []byte) bool {
	return bytes.Equal(pass1, pass2)
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
