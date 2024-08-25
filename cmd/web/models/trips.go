package models

import "trip-planner/internal/database"

func NewTripsModel(email string, trips *[]database.Trip) *TripsModel {
	return &TripsModel{NewMainLayout(email), trips}
}

type TripsModel struct {
	*MainLayoutModel
	Trips *[]database.Trip
}
