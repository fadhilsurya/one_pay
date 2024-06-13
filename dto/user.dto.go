package dto

type LoginDTO struct {
	PhoneNumber string `json:"phone_number" valid:"required,phone_number"`
	Password    string `json:"password" valid:"required"`
}

type RegisterDTO struct {
	FullName    string `json:"full_name" valid:"required"`
	Username    string `json:"username" valid:"required"`
	Address     string `json:"address" valid:"required"`
	Role        string `json:"role" valid:"required"`
	PhoneNumber string `json:"phone_number" valid:"required"`
	Password    string `json:"password" valid:"required"`
}
