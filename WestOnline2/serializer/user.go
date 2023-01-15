package serializer

import "WestOnline2/model"

type User struct {
	ID       uint   `json:"id" form:"id" `
	UserName string `json:"user_name" form:"user_name" `
	CreateAt int64  `json:"create_at" form:"create_at"`
}

func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}
