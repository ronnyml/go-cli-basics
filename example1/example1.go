package main

import (
    "fmt"
    "os"
)

func main() {
	cmd := os.Args[0]
	fmt.Printf("Program name: %s\n", cmd)
}