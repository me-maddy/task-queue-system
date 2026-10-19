package database

import (
	"time"
)

type User struct {
	ID          int64
	Name        string
	Email       string
	Phonenumber string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UpdateUserInput struct {
	Name        *string `json:"name"`
	Email       *string `json:"email"`
	Phonenumber *string `json:"phonenumber"`
	Password    *string `json:"password"`
}

type CreateUserInput struct {
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Email       string `json:"email" validate:"required,email"`
	Phonenumber string `json:"phonenumber" validate:"required,len=10"`
	Password    string `json:"password" validate:"required,len=6"`
}

var users = []User{}

func CreateUser(user User) User {
	users = append(users, user)
	return user
}

func UpdateUser(id int64, user UpdateUserInput) bool {
	for idx, usr := range users {
		if usr.ID == id {
			if user.Name != nil {
				users[idx].Name = *user.Name
			}
			if user.Email != nil {
				users[idx].Email = *user.Email
			}
			if user.Phonenumber != nil {
				users[idx].Phonenumber = *user.Phonenumber
			}
			if user.Password != nil {
				users[idx].Password = *user.Password
			}
			users[idx].UpdatedAt = time.Now()
			return true
		}
	}
	return false
}

func DeleteUser(id int64) bool {
	for idx, usr := range users {
		if usr.ID == id {
			users = append(users[:idx], users[idx+1:]...)
			return true
		}
	}
	return false
}

func GetUserByID(id int64) *User {
	for i := range users {
		if users[i].ID == id {
			return &users[i]
		}
	}
	return nil
}

func GetAllUsers() []User {
	return users
}

func GetUserByEmail(email string) *User {
	for i := range users {
		if users[i].Email == email {
			return &users[i]
		}
	}
	return nil
}
