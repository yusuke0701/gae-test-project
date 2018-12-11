package model

import "fmt"

// User は、ユーザーを表す
type User struct {
	Email  string
	Name   string
	Detail UserDetail
}

// UserDetail は、ユーザーの詳細情報を表す
type UserDetail struct {
	FirstName string
	LastName  string
	Address   string
}

// NewUser は、ユーザーを作る
func NewUser(email, firstName, lastName, address string) *User {
	return &User{
		Email: email,
		Name:  fmt.Sprintf("%s %s", firstName, lastName),
		Detail: UserDetail{
			FirstName: firstName,
			LastName:  lastName,
			Address:   address,
		},
	}
}
