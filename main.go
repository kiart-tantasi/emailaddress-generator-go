package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const USAGE_MESSAGE = "Usage: go run main.go <count, required> <domain, optional>"

func main() {
	// Check if a command-line argument was provided
	if len(os.Args) < 2 {
		fmt.Println(USAGE_MESSAGE)
		os.Exit(1)
	} else if os.Args[1] == "help" {
		fmt.Println(USAGE_MESSAGE)
		os.Exit(0)
	}

	// Parse the first argument as an integer
	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Error with invalid count argument: %s\n", err)
		os.Exit(1)
	}

	// Create or overwrite the file
	fileName := fmt.Sprintf("contacts%d.csv", count)
	file, err := os.Create(fileName)
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

	// Define domain based on argument indexed 2
	domain := ""
	if len(os.Args) > 2 && os.Args[2] != "" {
		domain = os.Args[2]
	} else {
		domain = fmt.Sprintf("loadtest%d.com", count)
	}
	fmt.Printf("Will use domain %s\n", domain)

	// Use buffered writer for performance
	for i := 1; i <= count; i++ {
		_, err := writer.WriteString(fmt.Sprintf("user%d@%s\n", i, domain))
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

	fmt.Println("Done! Wrote", count, "lines to", fileName)
}
