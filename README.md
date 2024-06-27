# Simple Go FileCrypt

This is a simple command-line tool written in Go for encrypting and decrypting files using the filecrypt library. The tool supports basic commands to encrypt and decrypt files with password protection.
## Installation
Install Go: Make sure you have Go installed on your machine. You can download it from the official Go website.
- Clone the repository:
```sh
    git clone https://github.com/Indraswara/simple-go-filecrypt.git
    cd simple-go-filecrypt
```
- Get dependencies:
```sh
    go get github.com/Indraswara/file-encrypt/filecrypt
    go get golang.org/x/term
```
## Usage

Run the tool using the following commands:
- Encrypt a File To encrypt a file, run:
```sh
    go run . encrypt /path-to-file
```
You will be prompted to enter and confirm a password. The file will be encrypted using the specified password.
- Decrypt a File To decrypt a file, run:

``` sh
    go run . decrypt /path-to-file
```

You will be prompted to enter the password used during encryption. The file will be decrypted using the specified password.
To see the help message, run:

```sh

go run . help
```
## Commands
- encrypt: Encrypts a file.
- decrypt: Decrypts a file.
- help: Prints the help message.

Example
Encrypting a File

```sh
    go run . encrypt example.txt
```

Decrypting a File

```sh
    go run . decrypt example.txt
```