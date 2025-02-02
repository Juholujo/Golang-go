package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

func readBytes(r io.Reader) error {
	buf := make([]byte, 1024*1024) // 1MB буфер
	start := time.Now()
	var totalRead int64

	for {
		n, err := r.Read(buf)
		if n > 0 {
			totalRead += int64(n)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	fmt.Printf("Read %d bytes from io.Reader in %s\n", totalRead, time.Since(start))
	return nil
}

func main() {
	// Замер начальной памяти
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Initial Memory Usage: %v MB\n", m.Alloc/1024/1024)

	// 2GB данных
	payload := bytes.Repeat([]byte("A"), 10*1024*1024*1024) // 2GB данных

	// Преобразуем payload в io.Reader
	start := time.Now()
	err := readBytes(bytes.NewReader(payload))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Execution Time: %s\n", time.Since(start))

	// Замер финальной памяти
	runtime.ReadMemStats(&m)
	fmt.Printf("Final Memory Usage: %v MB\n", m.Alloc/1024/1024)
}
