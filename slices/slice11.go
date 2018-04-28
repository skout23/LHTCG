package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	// build full "[N]array" which can be ref by a slice via array[:]
	var nonce [24]byte
	fmt.Println(nonce)
	io.ReadFull(rand.Reader, nonce[:])
	fmt.Println(nonce)
}
