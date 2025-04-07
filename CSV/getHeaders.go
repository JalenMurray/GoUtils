package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func printHeaders(files []string) {
	if len(files) == 0 {
		fmt.Println("⚠️ Please provide at least one CSV filename")
		return
	}

	for _, file := range files {
		fmt.Printf("📄 File: %s\n", filepath.Base(file))
		headers, err := getCSVHeaders(file)
		if err != nil {
			fmt.Printf("\t❌ Error reading headers: %v\n\n", err)
			continue
		}
		fmt.Println(" Headers:")
		for i, h := range headers {
			fmt.Printf("\t[%d] %s\n", i+1, h)
		}
		fmt.Println()
	}
}

func getCSVHeaders(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		return nil, err
	}

	return headers, nil
}
