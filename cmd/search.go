/*
Copyright © 2025 Harikesh Ranjan Sinha <ranjansinhaharikesh@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search [directory_path] [filename]",
	Short: "Search for a file by name in the directory",
	Long:  `Recursively searches for a file in the given directory and prints its path(s)`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]
		filename := args[1]

		err := searchFile(dir, filename)
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func searchFile(root, target string) error {
	found := false

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Name() == target {
			fmt.Printf("🔍 Found: %s\n", path)
			found = true
		}
		return nil
	})

	if err != nil {
		return err
	}

	if !found {
		fmt.Println("❌ No matching file found.")
	}

	return nil
}
