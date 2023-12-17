package cmd

import (
	"klijentske-tehnologije/migration"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Table migration",
	Long:  "Table migration",
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

func migrate() {
	migration.Migrate()
}
