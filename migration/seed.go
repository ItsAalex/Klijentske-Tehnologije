package migration

import (
	"klijentske-tehnologije/configs"
	"klijentske-tehnologije/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db, err := configs.Connect()
	user := models.User{Firstname: "Random name", Lastname: "Random Lastname", Email: "random@email.com",
		StudyProgram: "SRT", Index: "SEr 58/20", Comment: "comment"}
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

	question := models.Question{Name: "John", Surname: "Doe", Email: "john.doe@example.com", Question: "What is the meaning of life?"}
	db.Create(&question)
	log.Printf("Question created successfully")
}
