package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Error occurred")
		os.Exit(1)
	}
	val := os.Args[1]
	str := os.Args[2]

	switch val {
	case "sha256":
		fmt.Printf("%x \n", s256(str))
	case "sha384":
		fmt.Printf("%x \n", s384(str))
	case "sha512":
		fmt.Printf("%x \n", s512(str))
	default:
		fmt.Printf("%x \n", s256("str"))
	}
}

func s256(str string) [32]uint8 {
	return sha256.Sum256([]byte(str))
}

func s384(str string) [48]uint8 {
	return sha512.Sum384([]byte(str))
}

func s512(str string) [64]uint8 {
	return sha512.Sum512([]byte(str))
}
