package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

func readBytes(r io.Reader) error {

	buf := new(bytes.Buffer)
	readBytes, err := buf.ReadFrom(r)
	if err != nil {
		return err
	}
	fmt.Printf("Read %d bytes from io.Reader\n", readBytes)
	return nil
}

func main() {
	// 10GB данных
	payload := bytes.Repeat([]byte("A"), 10*1024*1024*1024) // 10GB данных

	// Преобразуем payload в io.Reader с помощью bytes.NewReader
	err := readBytes(bytes.NewReader(payload))
	if err != nil {
		log.Fatal(err)
	}
}
