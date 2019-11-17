package main

import (
	// "bytes"
	"fmt"
	"strings"
)

func main() {
	// this is bytes.buffer
	// var buffer bytes.Buffer
	// buffer.Write([]byte("bytes.Buffer example\n"))

	// this is strings.Builder
	var buffer strings.Builder
	buffer.Write([]byte("bytes.Buffer example\n"))

	fmt.Println(buffer.String())
}
