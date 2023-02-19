package models

type User struct {
	// gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
}

type LoginDetails struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
