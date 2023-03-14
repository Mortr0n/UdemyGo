package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	arg := os.Args[1]

	file, err := os.Open(arg)
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
