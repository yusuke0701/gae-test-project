package model

// Account は、一件のアカウント情報を表す
type Account struct {
	ID       string `json:"id"`
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	NickName string `json:"nick_name"`
	Password string `json:"password" binding:"required"`

	Meta
}

// LoginAccount は、ログインする際のアカウント情報を表す
type LoginAccount struct {
	ID       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
