package enum

type GameStateEnum uint8

const (
	Initialized GameStateEnum = iota + 1
	Running
	Ended
)

func (gs GameStateEnum) IsValid() bool {
	return Initialized <= gs && gs <= Ended
}
