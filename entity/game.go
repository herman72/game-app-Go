package entity

type Game struct {
	ID int
	Category string
	Questions []Question
	Players []User
}