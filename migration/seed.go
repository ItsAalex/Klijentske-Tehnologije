package migration

import (
	"klijentske-tehnologije/configs"
	"klijentske-tehnologije/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db, err := configs.Connect()

	hashedUserPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), 12)
	if err != nil {
		log.Fatal("hash password error")
		return
	}
	user := models.User{FirstName: "Random name", LastName: "Random Lastname", Email: "random@email.com", Password: string(hashedUserPassword),
		Address: "palilula", City: "Nis", Postcode: "18000", Phone: "069432224"}
	db.Create(&user)
	log.Printf("User created successfully with email address %s \n", user.Email)

	hashedAdminPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), 12)
	if err != nil {
		log.Fatal("hash password error")
		return
	}

	admin := models.Admin{Name: "Random Admin Name", Email: "random01@gmail.com", Password: string(hashedAdminPassword)}
	db.Create(&admin)

	log.Printf("Admin created successfully")
}
