package requests

type CreateUserRequest struct {
	Email     string `json:"Email" binding:"required,email" gorm:"coloumn:email;unique"`
	Password  string `json:"Password" binding:"required,min=8"`
	FirstName string `json:"FirstName" binding:"required,min=2,max=50"`
	LastName  string `json:"LastName" binding:"required,min=2,max=50"`
	Address   string `json:"Address" binding:"required"`
	City      string `json:"City" binding:"required"`
	Postcode  string `json:"Postcode" binding:"required"`
	Phone     string `json:"Phone" binding:"required"`
}
