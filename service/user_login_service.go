package service

import (
	"gim/model"
	"gim/serializer"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	Account string `form:"account" json:"account" binding:"required,min=5,max=30"`
	Pwd     string `form:"pwd" json:"pwd" binding:"required,min=6,max=16"`
}

// Login 用户登录函数
func (service *UserLoginService) Login() (model.User, *serializer.Response) {
	var user model.User

	if err := model.DB.Where("account = ?", service.Account).First(&user).Error; err != nil {
		return user, &serializer.Response{
			Status: 40001,
			Msg:    "账号或密码错误",
		}
	}

	if user.CheckPassword(service.Pwd) == false {
		return user, &serializer.Response{
			Status: 40001,
			Msg:    "账号或密码错误",
		}
	}
	return user, nil
}
