package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func readBytes(r []byte) error {
	start := time.Now()
	fmt.Printf("Read %d bytes from []byte in %s\n", len(r), time.Since(start))
	return nil
}

func main() {
	// Замер начальной памяти
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Initial Memory Usage: %v MB\n", m.Alloc/1024/1024)

	payload := make([]byte, 10*1024*1024*1024) // 10GB данных

	start := time.Now()
	err := readBytes(payload)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Execution Time: %s\n", time.Since(start))

	// Замер финальной памяти
	runtime.ReadMemStats(&m)
	fmt.Printf("Final Memory Usage: %v MB\n", m.Alloc/1024/1024)
}
