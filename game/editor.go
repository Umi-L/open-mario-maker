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
		newTile := objects.MakeObjectFromId(e.game.SelectedObject)

		mousePosX, mousePosY := ebiten.CursorPosition()

		tileObject := newTile.GetObject()

		tileObject.X = float32(mousePosX)
		tileObject.Y = float32(mousePosY)

		e.Objects = append(e.Objects, newTile)
	}

	for _, o := range e.Objects {
		e.game.drawStack.Add(o.Draw(), o.GetObject().ZIndex)
	}
}

func MakeEditor(game *Game) Editor {
	return Editor{game: game}
}
