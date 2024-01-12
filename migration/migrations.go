package migration

import (
	fmt "fmt"
	configs "klijentske-tehnologije/configs"
	models "klijentske-tehnologije/models"
)

func Migrate() {
	db, err := configs.Connect()
	// Check for errors when connecting to the database
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("Error getting SQL database: %v\n", err)
		return
	}
	fmt.Println("Before AutoMigrate Admin")
	db.AutoMigrate(&models.Admin{})
	fmt.Println("After AutoMigrate Admin")

	fmt.Println("Before AutoMigrate User")
	db.AutoMigrate(&models.User{})
	fmt.Println("After AutoMigrate User")

	defer sqlDB.Close()
}
