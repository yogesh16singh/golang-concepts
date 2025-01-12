package main

import (
	"fmt"
	"os"
)

type Report struct {
	Title string
	Body  string
}

func (r *Report) GenerateReport() string {
	return fmt.Sprintf("Title: %s\nBody: %s", r.Title, r.Body)
}

type FileManager struct{}

func (f *FileManager) SaveToFile(filename, data string) error {
	return os.WriteFile(filename, []byte(data), 0644)
}

func main() {
	report := Report{Title: "Monthly Report", Body: "Details about the month..."}
	fileManager := FileManager{}

	// Generate the report
	reportContent := report.GenerateReport()

	// Save the report
	err := fileManager.SaveToFile("report.txt", reportContent)
	if err != nil {
		fmt.Println("Error saving report:", err)
	}
}
