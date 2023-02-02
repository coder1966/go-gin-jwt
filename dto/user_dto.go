package dto

import "go-gin-jwt/model"

type UserDto struct {
	Name string
}

func ToUserDto(user model.User) UserDto {
	return UserDto{Name: user.LoginName}
}
