package entity


type Question struct {
	ID int
	Question string
	PossibleAnswer []string
	CorrectAnswer string
	Difficulty string
	Category string
}