package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	//!Original or long way to log out the response body for educational purposes
	// making a byte slice to use for filling with the response body from our request
	// bs := make([]byte, 99999)
	// // Body implements the io.Reader interface and needs us to provide a byte slice that it can fill with data
	// resp.Body.Read(bs)
	// // remember a byte slice is basically a string so an easy type conversion exists
	// fmt.Println(string(bs))
	//! End of Long way.
	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

// Making our own custom implementation of the write interface with the logWriter struct from above
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
