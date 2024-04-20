/*
------------------------------------------------------------------------------------------------------------------------

	Aşağıdaki demo örnekte marshalling yapılmıştır

------------------------------------------------------------------------------------------------------------------------
*/
package main

import (
	"SampleGoLand/csd/console"
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"-"` //transient
	Name     string `json:"name"`
}

func NewUser(username, password, name string) *User {
	return &User{Username: username, Password: password, Name: name}
}

func main() {
	for {
		username := console.ReadString("Input username:")

		if username == "" {
			break
		}

		password := console.ReadString("Input password:")
		name := console.ReadString("Input name:")
		user := NewUser(username, password, name)

		d, e := json.Marshal(&user)

		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "problem occurred:%s\n", e.Error())
		} else {
			fmt.Printf("Data:%s\n", d)
		}
	}
}
