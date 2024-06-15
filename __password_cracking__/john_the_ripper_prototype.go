package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func sha256_hash(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

func crack(hash string, word_list []string) (string, bool) {
	for _, word := range word_list {
		if sha256_hash(word) == hash {
			return word, true
		}
	}
	return "", false
}

func main() {
	__hash_to_crack__ := "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a362c454a4bac4d"
	__word_list__ := []string{"password", "123456", "123456789", "12345678", "12345", "1234567", "1234567890", "password1", "1234", "1234567", "12345678", "123456789", "1234567890", "123456", "password", "password1"}

	if password, found := crack(__hash_to_crack__, __word_list__); found {
		fmt.Printf("Password found: %s", password)	
	}else {
		fmt.Printf("Password not found")
	}
}