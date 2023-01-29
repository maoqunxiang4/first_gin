package dto

import "goPro/model"

type UserDTO struct {
	UserName string `json:"username"`
	Phone    string `json:"phone"`
}

func toUserDTO(user model.User) UserDTO {
	return UserDTO{
		UserName: user.UserName,
		Phone:    user.Phone,
	}
}
