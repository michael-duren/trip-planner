package models

func NewTripsModel(email string) *TripsModel {
	return &TripsModel{NewMainLayout(email)}
}

type TripsModel struct {
	*MainLayoutModel
}
