package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	file, err := os.Create("text.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.Name = "text.txt"
	io.WriteString(writer, "gzip.Writer example\n")
	writer.Close()
}
