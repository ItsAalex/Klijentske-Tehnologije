package cmd

import (
	"klijentske-tehnologije/migration"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Database seeder",
	Long:  "Database seeder for user and product",
	Run: func(cmd *cobra.Command, args []string) {
		seed()
	},
}

func seed() {
	migration.Seed()
}
