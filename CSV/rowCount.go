package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func printRowCount(files []string) {
	if len(files) == 0 {
		fmt.Println("⚠️ Please provide at least one CSV filename")
		return
	}

	for _, file := range files {
		fmt.Printf("📄 File: %s\n", filepath.Base(file))
		rows, err := countCSVRows(file)
		if err != nil {
			fmt.Printf("\t❌ Error counting rows: %v\n\n", err)
			continue
		}
		fmt.Printf("\tRows (excluding header): %d\n\n", rows)
	}
}

func countCSVRows(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	if _, err := reader.Read(); err != nil {
		return 0, fmt.Errorf("unable to read header: %w", err)
	}

	rowCount := 0
	for {
		_, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return rowCount, fmt.Errorf("error reading row: %w", err)
		}
		rowCount++
	}

	return rowCount, nil
}
