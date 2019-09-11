package model

import "errors"

type User struct {
	ID int
	Name string
	Role string
}

type Users [] User

func (u Users) Exists (id int) bool  {
	exists := false
	for _,user := range u{
		if user.ID == id{
			return true
		}
	}
	return exists
}

func (u Users) FindByName(name string)(User,error)  {
	for _,user := range u{
		if user.Name == name{
			return user,nil
		}
	}
	return User{},errors.New("User not_FOUND")
}