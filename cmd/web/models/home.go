package models

import "trip-planner/internal/auth"

func NewHomeModel(user *auth.UserDto) *HomeModel {
	return &HomeModel{NewMainLayout(user)}
}

type HomeModel struct {
	*MainLayoutModel
}
