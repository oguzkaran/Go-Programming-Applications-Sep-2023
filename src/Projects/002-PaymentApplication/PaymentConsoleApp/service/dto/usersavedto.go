package dto

import "time"

type UserSaveDTO struct {
	Username, Password, Name, Phone string
	BirthDate                       time.Time
}

func NewUserSaveDTO(username, password, name, phone, birthDateStr string) *UserSaveDTO {
	birthDate, err := time.Parse("02/01/2006", birthDateStr)

	if err != nil {
		return nil
	}

	return &UserSaveDTO{username, password, name, phone, birthDate}
}

//...
