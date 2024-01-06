package dto

type UserDTO struct {
	Username, Name, Phone string
}

func NewUserDTO(username, name, phone string) *UserDTO {
	return &UserDTO{username, name, phone}
}

//...
