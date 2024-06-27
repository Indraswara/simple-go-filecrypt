package main

import (
	"fmt"
	"os"
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
	fmt.Println("Encrypting file")
}

func decryptHandler() {
	fmt.Println("Decrypting file")
}

func getPassword() {

}

func validatePassword() {

}
