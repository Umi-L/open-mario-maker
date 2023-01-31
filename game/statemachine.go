package game

func RunStateMachine(state GameState, game *Game) {
	switch state {
	case MainMenu:
	case Editing:
		EditorUpdate(game)
	}
}
