package service

import (
	"gim/model"
	"gim/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Name       string `form:"name" json:"name" binding:"required,min=2,max=30"`
	Account    string `form:"account" json:"account" binding:"required,min=5,max=30"`
	Pwd        string `form:"pwd" json:"pwd" binding:"required,min=4,max=16"`
	PwdConfirm string `form:"pwd_confirm" json:"pwd_confirm" binding:"required,min=4,max=16"`
	Avatar     string `form:"avatar" json:"avatar" binding:""`
}

// Valid 验证表单
func (service *UserRegisterService) Valid() string {

	if service.Pwd != service.PwdConfirm {
		return "两次密码不一致"
	}

	count := 0
	model.DB.Model(&model.User{}).Where("name = ?", service.Name).Count(&count)
	if count > 0 {
		return "昵称被占用"
	}

	count = 0
	model.DB.Model(&model.User{}).Where("account = ?", service.Account).Count(&count)
	if count > 0 {
		return "账号已经注册"
	}

	return ""
}

// Register 用户注册
func (service *UserRegisterService) Register() (model.User, *serializer.Response) {
	user := model.User{
		Name:    service.Name,
		Account: service.Account,
		Avatar:  service.Avatar,
		Status:  model.Active,
	}

	// 表单验证
	if err := service.Valid(); err != "" {
		return user, &serializer.Response{
			Status: 40002,
			Msg:    err,
		}
	}

	// 加密密码
	if err := user.SetPassword(service.Pwd); err != nil {
		return user, &serializer.Response{
			Status: 40002,
			Msg:    "密码加密失败",
		}
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return user, &serializer.Response{
			Status: 40002,
			Msg:    "注册失败",
		}
	}

	return user, nil
}
