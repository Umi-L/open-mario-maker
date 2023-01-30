package gamestates

type GameState int32

const (
	MainMenu GameState = 0
	Playing            = 1
	Editing            = 2
)
