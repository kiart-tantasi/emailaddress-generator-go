package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const USAGE_MESSAGE = "Usage: go run main.go <count, required> <domain, optional> <output-filename, optional> <offset, optional> <prefix, optional>"

func main() {
	// ========================== ARGUMENTS ==========================
	// Basic checks
	if len(os.Args) < 2 {
		fmt.Println(USAGE_MESSAGE)
		os.Exit(1)
	}
	// help
	if os.Args[1] == "help" {
		fmt.Println(USAGE_MESSAGE)
		os.Exit(0)
	}
	// 1st argument: count
	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Error with invalid count argument: %s\n", err)
		os.Exit(1)
	}
	// 2nd argument: email domain
	domain := ""
	if len(os.Args) > 2 && os.Args[2] != "" {
		domain = os.Args[2]
	} else {
		domain = fmt.Sprintf("loadtest%d.com", count)
	}
	fmt.Printf("Will use domain %s\n", domain)
	// 3rd argument: output filename
	outputFilename := ""
	if len(os.Args) > 3 && os.Args[3] != "" {
		outputFilename = os.Args[3]
	} else {
		outputFilename = fmt.Sprintf("output%d.csv", count)
	}
	// 4th argument: offset
	offset := 1
	if len(os.Args) > 4 && os.Args[4] != "" {
		offset, err = strconv.Atoi(os.Args[4])
		if err != nil {
			fmt.Printf("Error with invalid offset: %s\n", err)
			os.Exit(1)
		}
	}
	// 5th argument: prefix
	prefix := ""
	if len(os.Args) > 5 && os.Args[5] != "" {
		prefix = os.Args[5]
	}
	// ========================== ARGUMENTS ==========================

	// Create output file
	file, err := os.Create(outputFilename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Writer
	writer := bufio.NewWriter(file)

	// CSV header
	_, err = writer.WriteString("email\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Use buffered writer for performance
	for i := offset; i <= count; i++ {
		_, err := writer.WriteString(fmt.Sprintf("%suser%d@%s\n", prefix, i, domain))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	// Flush remaining data to disk
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	fmt.Println("Done! Wrote", count, "email addresses to", outputFilename)
}
