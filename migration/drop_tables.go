package migration

import (
	"fmt"
	"klijentske-tehnologije/configs"
	"klijentske-tehnologije/models"
)

// DropTables function drops all tables in the database.
func DropTables() {
	db, err := configs.Connect()
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("Error getting SQL database: %v\n", err)
		return
	}
	defer sqlDB.Close()
	// Drop tables in reverse order to avoid foreign key constraints issues

	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Admin{})
	db.Migrator().DropTable(&models.Question{})
	fmt.Println("Tables dropped.")
}
