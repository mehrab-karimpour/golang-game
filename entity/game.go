package entity

type Game struct {
	ID          uint
	CategoryId  uint
	Title       string
	Description string
	Users       []User
	Questions   []Question
}
