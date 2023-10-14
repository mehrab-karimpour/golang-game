package entity

type Question struct {
	ID              uint
	GameID          uint
	CorrectAnswerID uint // correct answer id !
	Title           string
	Difficulty      string
	Answers         []Answer
}
