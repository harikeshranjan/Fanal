/*
Copyright © 2025 Harikesh Ranjan Sinha <ranjansinhaharikesh@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var topN int

type FileInfo struct {
	Name        string
	Path        string
	Size        int64
	ModTime     time.Time
	Extension   string
	WordCount   int
	LetterCount int
	LineCount   int
}

// topCmd represents the top command
var topCmd = &cobra.Command{
	Use:   "top [directory_path] --n=<N>",
	Short: "This command helps you to find top `N` files in the directory specified",
	Long:  `The keyword 'top' is used to find the top N files in the directory specified and prints the summary of the files by specifying size, extension, last modified, and words count, letters count and number of lines in the file`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dirPath := args[0]

		files, err := getTopFiles(dirPath, topN)
		if err != nil {
			fmt.Println("❌ Error:", err)
			return
		}

		fmt.Printf("📁 Top %d largest files in %s:\n\n", len(files), dirPath)

		for _, file := range files {
			printFileSummary(file)
		}
	},
}

func init() {
	rootCmd.AddCommand(topCmd)
	topCmd.Flags().IntVarP(&topN, "n", "n", 3, "Number of top files to display")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// topCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// topCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getTopFiles(dirPath string, topN int) ([]FileInfo, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	var files []FileInfo
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		fullPath := filepath.Join(dirPath, entry.Name())
		info, err := analyzeFile(fullPath)
		if err == nil {
			files = append(files, info)
		}
	}

	// Sort by size descending
	sort.Slice(files, func(i, j int) bool {
		return files[i].Size > files[j].Size
	})

	// Trim to top N
	if topN > len(files) {
		topN = len(files)
	}
	return files[:topN], nil
}

// analyzeFile returns FileInfo summary of a file
func analyzeFile(path string) (FileInfo, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return FileInfo{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return FileInfo{}, err
	}

	content := string(data)
	words := strings.Fields(content)
	lines := strings.Split(content, "\n")

	letterCount := 0
	for _, ch := range content {
		if ch != ' ' && ch != '\n' && ch != '\r' && ch != '\t' {
			letterCount++
		}
	}

	return FileInfo{
		Name:        filepath.Base(path),
		Path:        path,
		Size:        stat.Size(),
		ModTime:     stat.ModTime(),
		Extension:   filepath.Ext(path),
		WordCount:   len(words),
		LetterCount: letterCount,
		LineCount:   len(lines),
	}, nil
}

func printFileSummary(f FileInfo) {
	fmt.Printf("📄 %s\n", f.Name)
	fmt.Printf("  📏 Size: %d bytes\n", f.Size)
	fmt.Printf("  🧩 Extension: %s\n", f.Extension)
	fmt.Printf("  🕒 Last modified: %s\n", f.ModTime.Format(time.RFC1123))
	fmt.Printf("  🔤 Word count: %d\n", f.WordCount)
	fmt.Printf("  🔡 Letter count: %d\n", f.LetterCount)
	fmt.Printf("  📃 Line count: %d\n", f.LineCount-1)
	fmt.Println("-------------------------------")
}
