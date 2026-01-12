package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	// Get current date in YYYYMMDD format
	today := time.Now()
	dateStr := today.Format("20060102")

	// Define template configurations
	templates := map[string]struct {
		templateFile string
		outputPrefix string
	}{
		"traverse": {"traverse_template.tmpl", "traverse"},
		"dag":      {"dag_template.tmpl", "dag"},
		"dijkstra": {"dijkstra_template.tmpl", "dijkstra"},
	}

	generatedCount := 0
	var generatedFiles []string

	// Generate all templates
	for templateType, config := range templates {
		// Create output filename
		outputFile := fmt.Sprintf("%s_%s.go", config.outputPrefix, dateStr)

		// Check if file already exists in current directory
		outputPath := outputFile
		if _, err := os.Stat(outputPath); err == nil {
			fmt.Printf("File %s already exists for today\n", outputFile)
			continue
		}

		// Read template file
		templateFile, err := os.Open(config.templateFile)
		if err != nil {
			fmt.Printf("⚠️  Warning: Could not open template file %s: %v\n", config.templateFile, err)
			continue
		}

		// Create output file
		outputFileHandle, err := os.Create(outputPath)
		if err != nil {
			templateFile.Close()
			fmt.Printf("⚠️  Warning: Could not create output file %s: %v\n", outputPath, err)
			continue
		}

		// Add header comment
		header := fmt.Sprintf("// Generated on %s\n// Daily practice file: %s (%s template)\n\n",
			today.Format("2006-01-02 15:04:05"), outputFile, templateType)
		outputFileHandle.WriteString(header)

		// Copy template content to output file
		_, err = io.Copy(outputFileHandle, templateFile)
		templateFile.Close()
		outputFileHandle.Close()

		if err != nil {
			fmt.Printf("⚠️  Warning: Error copying template to output file %s: %v\n", outputFile, err)
			continue
		}

		generatedFiles = append(generatedFiles, outputFile)
		generatedCount++
	}

	// Print summary
	if generatedCount == 0 {
		fmt.Println("No new files generated - all templates already exist for today")
	} else {
		fmt.Printf("✅ Successfully generated %d file(s) for today's practice:\n", generatedCount)
		for _, file := range generatedFiles {
			fmt.Printf("   - %s\n", file)
		}
	}
}
