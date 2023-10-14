package enum

type UserAnswerEnum uint8

const (
	A UserAnswerEnum = iota + 1
	B
	C
	D
)

func (ua UserAnswerEnum) IsValid() bool {
	return A <= ua && ua <= D
}
