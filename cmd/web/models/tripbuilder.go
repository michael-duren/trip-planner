package models

func NewTripBuildersModel(email, tripname string, tripid int) *TripBuildersModel {
	return &TripBuildersModel{NewMainLayout(email), tripname, tripid}
}

type TripBuildersModel struct {
	*MainLayoutModel
	Tripname string
	Tripid   int
}
