package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// Get current date in YYYYMMDD format
	today := time.Now()
	dateStr := today.Format("20060102")

	// Create output filename
	outputFile := fmt.Sprintf("sortalgo_%s.go", dateStr)

	// Check if file already exists in parent directory
	outputPath := outputFile
	if _, err := os.Stat(outputPath); err == nil {
		fmt.Printf("File %s already exists for today\n", outputFile)
		return
	}

	// Read template file
	templateFile, err := os.Open("sortalgo_template.tmpl")
	if err != nil {
		log.Fatalf("Error opening template file: %v", err)
	}
	defer templateFile.Close()

	// Create output file in parent directory
	outputFileHandle, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outputFileHandle.Close()

	// Add header comment
	header := fmt.Sprintf("// Generated on %s\n// Daily practice file: %s\n\n",
		today.Format("2006-01-02 15:04:05"), outputFile)
	outputFileHandle.WriteString(header)

	// Copy template content to output file
	_, err = io.Copy(outputFileHandle, templateFile)
	if err != nil {
		log.Fatalf("Error copying template to output file: %v", err)
	}

	fmt.Printf("✅ Successfully generated %s for today's practice!\n", outputFile)
}

