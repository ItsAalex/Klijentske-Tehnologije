package cmd

import (
	"klijentske-tehnologije/migration"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on dev server",
	Long:  "Application will be served on host and port defined in .env file",
	Run: func(cmd *cobra.Command, args []string) {
		migration.Serve()
	},
}
