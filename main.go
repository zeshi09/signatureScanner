package main

import (
	"fmt"
	"os"

	"github.com/zeshi09/signatureScanner/internal/scan"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Use: scanner <path>")
		os.Exit(1)
	}

	path := os.Args[1]
	findings := scan.Run(path)

	fmt.Printf("Found signatures: %d\n", len(findings))
}
