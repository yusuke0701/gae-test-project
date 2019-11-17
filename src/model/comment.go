package model

// Comment は、一件のコメントを表す
type Comment struct {
	ID   string `json:"id" binding:"required"`
	Body string `json:"body" binding:"required"`
}
