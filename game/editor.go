package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/umi-l/open-mario-maker/objects"
)

type Editor struct {
	game *Game

	Objects []objects.ObjectInterface
}

var leftMouseDown bool

func (e *Editor) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		leftMouseDown = true
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		leftMouseDown = false
	}

	if leftMouseDown {
		e.Objects = append(e.Objects, objects.MakeObjectFromId(e.game.SelectedObject))
	}

	for _, o := range e.Objects {

	}
}

func MakeEditor(game *Game) Editor {
	return Editor{game: game}
}
