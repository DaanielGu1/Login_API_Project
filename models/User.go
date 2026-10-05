package models

type User struct {	
	Username string
	First_name string
	Last_name string
	Birthday string
	Email string
	Password string
}

func NewUser(username string, f_name string, l_name string, birth string, email string, psw string) *User {
	return &User{Username: username, First_name: f_name, Last_name: l_name, Birthday: birth, Email: email, Password: psw}
}

