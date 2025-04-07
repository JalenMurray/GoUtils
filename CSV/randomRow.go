package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

func printRandomRow(files []string) {
	if len(files) == 0 {
		fmt.Println("⚠️ Please provide at least one CSV filename")
		return
	}

	for _, file := range files {
		fmt.Printf("📄 File: %s\n", filepath.Base(file))
		rowJSON, err := getRandomRowAsJSON(file)
		if err != nil {
			fmt.Printf("\t❌ Error reading headers: %v\n\n", err)
			continue
		}

		fmt.Println("\t🎲 Random Row:")
		fmt.Println("\t" + rowJSON)
		fmt.Println()
	}
}

func getRandomRowAsJSON(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	headers, err := reader.Read()
	if err != nil {
		return "", fmt.Errorf("error reading headers: %w", err)
	}

	rows, err := reader.ReadAll()
	if err != nil {
		return "", fmt.Errorf("error reading rows: %w", err)
	}
	if len(rows) == 0 {
		return "", fmt.Errorf("no data rows found")
	}

	randomRow := rows[rand.Intn(len(rows))]

	rowMap := make(map[string]string)
	for i, header := range headers {
		if i < len(randomRow) {
			rowMap[header] = randomRow[i]
		} else {
			rowMap[header] = ""
		}
	}

	jsonData, err := json.MarshalIndent(rowMap, "\t", "\t")
	if err != nil {
		return "", fmt.Errorf("error converting to JSON: %w", err)
	}

	return string(jsonData), nil
}
