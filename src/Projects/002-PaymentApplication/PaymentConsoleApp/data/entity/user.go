package entity

import "time"

type User struct {
	Username, Password, Name, Phone string
	BirthDate                       time.Time
}

func NewUser(username, password, name, phone string, birthDate time.Time) *User {
	return &User{username, password, name, phone, birthDate}
}

//...
