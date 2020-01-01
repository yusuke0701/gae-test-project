package model

// Comment は、一件のコメントを表す
type Comment struct {
	ID         int64  `json:"id"`
	ThreadID   string `json:"thread_id" binding:"required"`
	Body       string `json:"body" binding:"required"`
	ContentURL string `json:"content_url"`

	Meta
}
