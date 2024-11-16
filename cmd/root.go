package cmd

import (
	"book-store/cmd/api"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(api.Cmd)
}

var rootCmd = &cobra.Command{
	Use:   "book-store",
	Short: "book-store",
	Long:  "book-store",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
