package entity

type User struct {
	Username, Password, Name, Phone string
}

func NewUser(username, password, name, phone string) *User {
	return &User{username, password, name, phone}
}

//...
