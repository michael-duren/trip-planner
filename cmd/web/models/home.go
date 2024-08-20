package models

func NewHomeModel(email string) *HomeModel {
	return &HomeModel{NewMainLayout(email)}
}

type HomeModel struct {
	*MainLayoutModel
}
