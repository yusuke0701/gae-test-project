package model

import (
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/boom"
)

// Comment は、コメントを表す
type Comment struct {
	ParentKey datastore.Key `json:"-" datastore:"-" boom:"parent"`
	ID        int64         `json:"id" datastore:"-" boom:"id"`
	Title     string        `json:"title"`
	Body      string        `json:"body"`
}

// NewComment は、コメントを作る
func NewComment(b *boom.Boom, email string, id int64, title, body string) (*Comment, error) {
	user := NewUser(email, "", "", "")
	key, err := b.KeyError(user)
	if err != nil {
		return nil, err
	}
	return &Comment{
		ParentKey: key,
		ID:        id,
		Title:     title,
		Body:      body,
	}, nil
}
