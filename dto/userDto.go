package dto

import "github.com/bubbly-hb/blogSystem-gin-vue/model"

type UserDto struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func ToUserDto(user *model.User) *UserDto {
	return &UserDto{
		Name:  user.Name,
		Email: user.Email,
	}
}
