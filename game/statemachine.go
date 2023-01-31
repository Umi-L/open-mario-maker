package game

var editor Editor

func InitStateMachine(game *Game) {
	editor = MakeEditor(game)
}

func RunStateMachine(state GameState) {
	switch state {
	case MainMenu:
	case Editing:
		editor.Update()
	}
}
