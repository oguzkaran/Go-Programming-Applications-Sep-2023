package dto

type UserDTO struct {
	Username, Name, Phone string
}

func NewUser(username, name, phone string) *UserDTO {
	return &UserDTO{username, name, phone}
}

//...
