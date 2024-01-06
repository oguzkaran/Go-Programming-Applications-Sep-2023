package dto

type UserSaveDTO struct {
	Username, Password, Name, Phone string
}

func NewUserSaveDTO(username, password, name, phone string) *UserSaveDTO {
	return &UserSaveDTO{username, password, name, phone}
}

//...
