package serivice

import (
	"WestOnline2/model"
	"WestOnline2/pkg/utils"
	"WestOnline2/serializer"
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	gorm.Model
	UserName       string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" `
	PasswordDigest string `form:"password_digest" json:"password" binding:"required,min=5,max=16"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int

	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "已经有这个人了",
		}
	}
	user.UserName = service.UserName

	if err := user.SetPassword(service.PasswordDigest); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}

	if err := model.DB.Create(&user).Error; err != nil {
		fmt.Println(err)
		return serializer.Response{
			Status: 500,
			Msg:    "数据库操作失误",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "用户注册成功",
	}
}

func (service *UserService) Login() serializer.Response {
	var user model.User

	err := model.DB.Where("user_name=?", service.UserName).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "用户不存在，请先登录",
			}
		}
		return serializer.Response{
			Status: 500,
			Msg:    "数据库错误",
		}
	}
	if user.CheckPassword(service.PasswordDigest) == false {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}

	token, err := utils.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {

		return serializer.Response{
			Status: 500,
			Msg:    "Token签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    "登陆成功",
	}
}
