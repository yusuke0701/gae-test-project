package model

// Tag は、スレッドのタグを表す
type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name" binding:"required"`

	Meta
}
