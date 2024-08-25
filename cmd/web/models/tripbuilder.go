package models

import (
	"trip-planner/internal/auth"
	"trip-planner/internal/database"
)

func NewTripBuildersModel(user *auth.UserDto, trip database.Trip) *TripBuildersModel {
	return &TripBuildersModel{NewMainLayout(user), trip}
}

type TripBuildersModel struct {
	*MainLayoutModel
    Trip database.Trip
}
