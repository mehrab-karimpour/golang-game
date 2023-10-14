package entity

type Category struct {
	ID      uint
	Title   string
	Content string
	Games   []Game
}
