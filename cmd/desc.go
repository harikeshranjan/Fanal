/*
Copyright © 2025 Harikesh Ranjan Sinha <ranjansinhaharikesh@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// descCmd represents the desc command
var descCmd = &cobra.Command{
	Use:   "desc [directory_path]",
	Short: "Describe the content of a directory",
	Long:  `The desc command analyzes and prints basic information about the files in the given directory`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dirPath := args[0]
		err := describeDirectory(dirPath)
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(descCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// descCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// descCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func describeDirectory(dirPath string) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	absPath, _ := filepath.Abs(dirPath)
	dirName := filepath.Base(absPath)

	fmt.Println("═════════════════════════════════════════════════")
	fmt.Printf("📁 DIRECTORY: %s\n", dirName)
	fmt.Printf("📍 Path: %s\n", absPath)
	fmt.Println("═════════════════════════════════════════════════")
	fmt.Println()

	var totalFiles int
	var totalSize int64
	var fileTypes = make(map[string]int)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"File Name", "Size", "Extension", "Last Modified"})
	table.SetRowLine(true)

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

		table.Append([]string{name, size, ext, modified})

		totalFiles++
		totalSize += info.Size()
		fileTypes[ext]++
	}

	table.Render()

	// Print summary
	fmt.Println()
	fmt.Println("📊 Summary")
	fmt.Println("═════════════════════════════════════════════════")
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

	return nil
}

// formatFileSize formats file size in bytes to a more readable format
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
