package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/atotto/clipboard"
)

const (
	letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	letterIdxMask = 0x3F
)

var (
	number = flag.Int("n", 6, "The number of password")
)

func main() {
	flag.Parse()

	buffer := make([]byte, *number)
	if _, err := rand.Read(buffer); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	for i := range buffer {
		idx := int(buffer[i] & letterIdxMask)
		if idx < len(letters) {
			buffer[i] = letters[idx]
			i++
		} else {
			if _, err := rand.Read(buffer[i : i+1]); err != nil {
				log.Println(err)
				os.Exit(1)
			}
		}
	}

	pw := string(buffer)
	clipboard.WriteAll(string(pw))
	fmt.Println(pw, "\nfinish to create and copy password completely!")
}
