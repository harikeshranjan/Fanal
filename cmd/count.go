/*
Copyright © 2025 Harikesh Ranjan Sinha <ranjansinhaharikesh@gmail.com>
*/
package cmd

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var fileExt string

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count [directory] [substring]",
	Short: "Count files matching name or extension filters",
	Long:  `The count command traverses the given directory and counts files matching the substring and optional extension filter.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dirPath := args[0]
		substr := strings.ToLower(args[1])

		files, count, err := countMatchingFiles(dirPath, substr, fileExt)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("✅ Total matching files: %d\n", count)
		fmt.Println("Matched files:")
		for i, file := range files {
			fmt.Printf("[%d]. [%s]\n", i+1, file)
		}
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
	countCmd.Flags().StringVarP(&fileExt, "ext", "f", "", "filter by file extension (e.g., .go)")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// countCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// countCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func countMatchingFiles(dirPath, substr, ext string) ([]string, int, error) {
	var count int
	var files []string

	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		name := strings.ToLower(d.Name())

		if strings.Contains(name, substr) {
			if ext == "" || strings.HasSuffix(name, strings.ToLower(ext)) {
				count++
				files = append(files, path)
			}
		}

		return nil
	})

	return files, count, err
}
