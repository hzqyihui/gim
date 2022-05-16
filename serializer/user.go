package serializer

import "gim/model"

// User 用户序列化器
type User struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Account string `json:"account"`
	Status  int    `json:"status"`
	Avatar  string `json:"avatar"`
}

// UserResponse 单个用户序列化
type UserResponse struct {
	Response
	Data User `json:"data"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:      user.ID,
		Name:    user.Name,
		Account: user.Account,
		Status:  user.Status,
		Avatar:  user.Avatar,
	}
}

//todo 这里感觉每个结构体都会涉及到构建响应体，后续优化为 反射类型
// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) UserResponse {
	return UserResponse{
		Data: BuildUser(user),
	}
}
