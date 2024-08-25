package models

import "trip-planner/internal/auth"

func NewMainLayout(user *auth.UserDto) *MainLayoutModel {
	return &MainLayoutModel{user}
}

type MainLayoutModel struct {
	User *auth.UserDto
}
