package main

import "fmt"

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcsv-utils.exe <command> [args]")
	fmt.Println("\nAvailable commands:")
	fmt.Println("\tget-headers <file1> <file2> ... <fileN> Show headers of provided CSV files")
	fmt.Println("\tget-rows <file1> <file2> ... <fileN> Show number of rows in provided CSV files")
	fmt.Println("\trandom-row <file1> <file2> ... <fileN> Show a random row from provided CSV files")
}
