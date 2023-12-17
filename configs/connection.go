package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

// Connection returns the global database instance
func Connection() *gorm.DB {
	return Db
}

// Connect connects to the database using the configuration settings
func Connect() (*gorm.DB, error) {
	// Load database configurations from the config file
	LoadConfig()

	// Get the loaded configurations
	cfg := Get()

	// Construct the Data Source Name (DSN) for the MySQL database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)

	// Open a connection to the MySQL database using the provided DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return nil, err
	}

	// Enable SQL query logging for debugging purposes
	db = db.Debug()

	// Return the connected database instance
	return db, nil
}
