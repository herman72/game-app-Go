package dto

// data transfer object for user registration
type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}


type RegisterResponse struct {
	User UserInfo `json:"user"`
			
}
