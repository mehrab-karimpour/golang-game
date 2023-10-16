package entity

type User struct {
	ID          uint
	FirstName   string
	LastName    string
	PhoneNumber string
	Password    string
	CreatedAt   string
	UpdatedAt   string
	Games       []Game
}
