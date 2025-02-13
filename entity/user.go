package entity


type User struct {
	ID uint
	PhoneNumber string
	Name string
	// password always keep hash password
	Password string
}