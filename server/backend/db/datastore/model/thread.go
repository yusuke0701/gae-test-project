package model

// Thread は、1つのスレッドを表す
type Thread struct {
	ID     string  `json:"id"`
	TagIDs []int64 `json:"tag_ids"`
	Title  string  `json:"title" binding:"required"`

	Meta
}
