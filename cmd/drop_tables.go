package cmd

import (
	migration "klijentske-tehnologije/migration"

	cobra "github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dropCmd)
}

var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop all Tables in Schema",
	Long:  "Drop all Tables in Schema on command go run main.go drop",
	Run: func(cmd *cobra.Command, args []string) {
		drop()
	},
}

func drop() {
	migration.DropTables()
}
