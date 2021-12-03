package model

type User struct {
	Account   string `json:"account" gorm:"column:account;primaryKey"`
	Password  string `json:"password" gorm:"uniqueIndex"`
	Email     string `json:"email"`
	NickName  string `json:"nick_name"`
	AvatarExt string `json:"avatar_ext"`
}

type UserMessage struct {
	Account  string `json:"account"`
	Email    string `json:"email"`
	NickName string `json:"nick_name"`
}
