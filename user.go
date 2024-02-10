package main

import "errors"

type User struct {
	Name     string
	Email    string
	Password string
}

var users = make(map[string]*User)

func (u *User) Save() error {

	if _, ok := users[u.Email]; ok {
		return errors.New("User already exists")
	}	

	users[u.Email] = u
	
	return nil
}
