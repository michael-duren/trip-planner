package models

func NewMainLayout(email string) *MainLayoutModel {
	return &MainLayoutModel{email}
}

type MainLayoutModel struct {
	Email string
}
