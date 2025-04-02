package main

import (
	"fmt"
	"main/pkg"
	"time"
    "os"
    "bufio"
)

// green := "\x1b[38;5;119m"
// white := "\x1b[37m"
// red := "\x1b[31m"
func openFileAndHash(namefile string) {
    file, err := os.Open(namefile)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    lineNumber := 1
    for scanner.Scan() {
        line := scanner.Text()
        hashed, err := pkg.Hash(line, pkg.CoffeeSalt())
        if err != nil {
            fmt.Printf("Error hashing password %d: %s\n", lineNumber, err)
            continue
        }
        fmt.Printf("Password \x1b[31m%d\x1b[0m: \x1b[38;5;46m%s\x1b[0m -> \x1b[33m%s\x1b[0m\n", lineNumber, line, hashed)
        lineNumber++
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
}



func main() {
	start := time.Now()
    openFileAndHash("password_list.txt")

	fmt.Println("Time:", time.Since(start))
}
