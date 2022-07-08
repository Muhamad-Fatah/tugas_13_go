package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main() {
	var nama = "Muhamad Fatah"

	// Base64
	var encode = base64.StdEncoding.EncodeToString([]byte(nama))
	fmt.Println(encode)

	// Hash SHA1
	var sha = sha1.New()
	sha.Write([]byte(encode))
	var enkripsi = sha.Sum(nil)
	var stringenkripsi = fmt.Sprintf("%x", enkripsi)

	fmt.Println(stringenkripsi)
}
