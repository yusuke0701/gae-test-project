package model

import "time"

// Comment は、一件のコメントを表す
type Comment struct {
	ID        string    `json:"id" binding:"required"`
	Body      string    `json:"body" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
