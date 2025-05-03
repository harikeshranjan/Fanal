/*
Copyright © 2025 Harikesh Ranjan Sinha <ranjansinhaharikesh@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// fsumCmd represents the fsum command
var fsumCmd = &cobra.Command{
	Use:   "fsum [directory_path] [filename]",
	Short: "This command helps you to find the summary of a file",
	Long:  `Using the fsum command you can easily get the summary of the whole file. The summary contains the word count, character count, size, extension of the file, number of lines of content in the file and last modified`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dirPath := args[0]
		filename := args[1]
		fullPath := filepath.Join(dirPath, filename)

		info, err := os.Stat(fullPath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		data, err := os.ReadFile(fullPath)
		if err != nil {
			fmt.Println("Error reading file:", err)
		}

		content := string(data)
		words := strings.Fields(content)
		lines := strings.Split(content, "\n")

		fmt.Println("📄 File Summary")
		fmt.Println("--------------")
		fmt.Printf("🔤 Word count: %d\n", len(words))
		fmt.Printf("🔡 Character count: %d\n", len(content))
		fmt.Printf("📏 File size: %d bytes\n", info.Size())
		fmt.Printf("🧩 Extension: %s\n", filepath.Ext(filename))
		fmt.Printf("📃 Line count: %d\n", len(lines))
		fmt.Printf("🕒 Last modified: %s\n", info.ModTime().Format(time.RFC1123))
	},
}

func init() {
	rootCmd.AddCommand(fsumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fsumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fsumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
