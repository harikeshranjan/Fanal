/*
Copyright © 2025 Harikesh Ranjan Sinha <ranjansinhaharikesh@gmail.com>
*/
package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// FileSummary holds file details for summary/export
type FileSummary struct {
	Name      string `json:"name"`
	Size      string `json:"size"`
	Extension string `json:"extension"`
	Modified  string `json:"modified"`
}

// descCmd represents the desc command
var descCmd = &cobra.Command{
	Use:   "desc [directory_path]",
	Short: "Describe the content of a directory",
	Long:  `The desc command analyzes and prints basic information about the files in the given directory`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dirPath := args[0]
		exportFormat, _ := cmd.Flags().GetString("export")

		summaries, totalFiles, totalSize, fileTypes, err := describeDirectory(dirPath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if exportFormat == "json" {
			exportAsJSON(summaries)
		} else if exportFormat == "csv" {
			exportAsCSV(summaries)
		} else {
			renderTable(summaries)
			printSummary(dirPath, totalFiles, totalSize, fileTypes)
		}
	},
}

func init() {
	rootCmd.AddCommand(descCmd)
	descCmd.Flags().String("export", "", "Export the result to a file (options: json, csv)")
}

// describeDirectory returns file summaries and basic stats
func describeDirectory(dirPath string) ([]FileSummary, int, int64, map[string]int, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, 0, 0, nil, err
	}

	var summaries []FileSummary
	var totalFiles int
	var totalSize int64
	fileTypes := make(map[string]int)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			continue
		}

		name := info.Name()
		size := formatFileSize(info.Size())
		ext := strings.ToLower(filepath.Ext(name))
		modified := info.ModTime().Format("02 Jan 06 03:04 PM")

		summaries = append(summaries, FileSummary{
			Name:      name,
			Size:      size,
			Extension: ext,
			Modified:  modified,
		})

		totalFiles++
		totalSize += info.Size()
		fileTypes[ext]++
	}
	return summaries, totalFiles, totalSize, fileTypes, nil
}

// renderTable displays the table in terminal
func renderTable(summaries []FileSummary) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"File Name", "Size", "Extension", "Last Modified"})
	table.SetRowLine(true)

	for _, s := range summaries {
		table.Append([]string{s.Name, s.Size, s.Extension, s.Modified})
	}
	table.Render()
}

// printSummary prints stats after the table
func printSummary(dirPath string, totalFiles int, totalSize int64, fileTypes map[string]int) {
	absPath, _ := filepath.Abs(dirPath)
	dirName := filepath.Base(absPath)

	fmt.Println()
	fmt.Println("📊 Summary")
	fmt.Println("═════════════════════════════════════════════════")
	fmt.Printf("📁 DIRECTORY: %s\n", dirName)
	fmt.Printf("📍 Path: %s\n", absPath)
	fmt.Printf("📄 Total Files: %d\n", totalFiles)
	fmt.Printf("📦 Total Size: %s\n", formatFileSize(totalSize))

	if len(fileTypes) > 0 {
		fmt.Println("📁 File Types:")
		for ext, count := range fileTypes {
			if ext == "" {
				fmt.Printf("   (no ext): %d\n", count)
			} else {
				fmt.Printf("   %s: %d\n", ext, count)
			}
		}
	}
	fmt.Println("═════════════════════════════════════════════════")
}

// formatFileSize formats file size for readability
func formatFileSize(size int64) string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
	)

	switch {
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/float64(GB))
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/float64(MB))
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/float64(KB))
	default:
		return fmt.Sprintf("%d B", size)
	}
}

// exportAsJSON saves data to desc_output.json
func exportAsJSON(data []FileSummary) {
	file, err := os.Create("desc_output.json")
	if err != nil {
		fmt.Println("❌ Could not create JSON file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		fmt.Println("❌ Failed to encode JSON:", err)
	} else {
		fmt.Println("✅ Exported to desc_output.json")
	}
}

// exportAsCSV saves data to desc_output.csv
func exportAsCSV(data []FileSummary) {
	file, err := os.Create("desc_output.csv")
	if err != nil {
		fmt.Println("❌ Could not create CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Name", "Size", "Extension", "Modified"})
	for _, d := range data {
		writer.Write([]string{d.Name, d.Size, d.Extension, d.Modified})
	}
	fmt.Println("✅ Exported to desc_output.csv")
}
