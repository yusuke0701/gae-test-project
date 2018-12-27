package model

import "fmt"

// User は、ユーザーを表す
type User struct {
	Email  string     `json:"email" datastore:"-" boom:"id"`
	Name   string     `json:"name"`
	Detail UserDetail `json:"detail" datastore:",flatten"`
}

// UserDetail は、ユーザーの詳細情報を表す
type UserDetail struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   string `json:"address"`
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
