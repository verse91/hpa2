//go:build ignore
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"bufio"
    "time"
)

func main() {
    start := time.Now()
	file, err := os.Open("password_list.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		hash := md5.Sum([]byte(line))
		hashString := hex.EncodeToString(hash[:])
		fmt.Printf("Password %d: %s -> %s\n", lineNumber, line, hashString)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
    fmt.Println("Time:", time.Since(start))
}
