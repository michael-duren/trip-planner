package models

import (
	"trip-planner/internal/auth"
	"trip-planner/internal/database"
)

func NewTripsModel(user *auth.UserDto, trips *[]database.Trip) *TripsModel {
	return &TripsModel{NewMainLayout(user), trips}
}

type TripsModel struct {
	*MainLayoutModel
	Trips *[]database.Trip
}
