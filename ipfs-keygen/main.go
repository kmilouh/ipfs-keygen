package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	// Generate key buffer
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil { // Error generating key
		log.Fatalln("While trying to read random source:", err)
	}
	// Output key swarm values
	fmt.Println("/key/swarm/psk/1.0.0/")
	fmt.Println("/base16/")
	fmt.Print(hex.EncodeToString(key))
}
