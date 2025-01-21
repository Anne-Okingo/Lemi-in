package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PrintFileContents(string) {
	// Print the input data
	content, err := ValidContent(os.Args[1])
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		return
	}
	// Print number of ants and room configuration
	fmt.Println(content[0])
	for _, line := range content[1:] {
		if !strings.Contains(line, "-") {
			fmt.Println(line)
		}
	}
	// Print links
	for _, line := range content[1:] {
		if strings.Contains(line, "-") {
			fmt.Println(line)
		}
	}
	fmt.Println()
}

func ValidContent(filename string) ([]string, error) {
	fileContent, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	defer fileContent.Close()

	validContent := []string{}
	scanner := bufio.NewScanner(fileContent)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		validContent = append(validContent, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %v", err)
	}

	return validContent, nil
}
